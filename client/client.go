package client

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	server     string
	authstr    string
	sessionkey string
	httpclient http.Client
}

func NewClient(server, authstr string) *Client {
	return &Client{server: server, authstr: authstr}
}

func (c *Client) Login() error {
	var status apiResponse
	var cookies []*http.Cookie

	c.httpclient = http.Client{}
	buff := []byte("/api/login/" + c.authstr)
	resp, err := c.httpclient.Post(c.server+"/api/", "", bytes.NewReader(buff))
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	xml.Unmarshal(b, &status)
	statusinfo, err := parseStatusInfo(status)
	if err != nil {
		return err
	}
	c.sessionkey = statusinfo.SessionKey

	u, err := url.Parse(c.server)
	if err != nil {
		return err
	}
	cookies = append(cookies, &http.Cookie{Name: "wbisessionkey", Value: c.sessionkey})
	cookies = append(cookies, &http.Cookie{Name: "wbiusername", Value: ""})
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	jar.SetCookies(u, cookies)
	c.httpclient = http.Client{Jar: jar}

	return nil
}

func (c Client) GetControllerStatistics() (map[string]ControllerStatistics, error) {
	xmlresponse, err := c.readXMLData("show/controller-statistics")
	if err != nil {
		return nil, err
	}
	return parseControllerStatistics(*xmlresponse)
}

func (c Client) GetHostPortStatistics() (map[string]HostPortStatistics, error) {
	xmlresponse, err := c.readXMLData("show/host-port-statistics")
	if err != nil {
		return nil, err
	}
	return parseHostPortStatistics(*xmlresponse)
}

func (c Client) GetDiskStatistics() (map[string]DiskStatistics, error) {
	xmlresponse, err := c.readXMLData("show/disk-statistics")
	if err != nil {
		return nil, err
	}
	stats, err := parseDiskStatistics(*xmlresponse)
	if err != nil {
		return nil, err
	}

	xmlresponse, err = c.readXMLData("show/disks")
	if err != nil {
		return nil, err
	}
	info, err := parseDiskInfo(*xmlresponse)
	if err != nil {
		return nil, err
	}

	for name, stat := range stats {
		stat.Health = info[name].Health
		stat.TotalDataTransferred = info[name].TotalDataTransferred
		stat.Location = info[name].Location
		stats[name] = stat
	}
	return stats, nil
}

func (c Client) GetVdiskStatistics() (map[string]VdiskStatistics, error) {
	xmlresponse, err := c.readXMLData("show/vdisk-statistics")
	if err != nil {
		return nil, err
	}
	stats, err := parseVdiskStatistics(*xmlresponse)
	if err != nil {
		return nil, err
	}

	xmlresponse, err = c.readXMLData("show/vdisks")
	if err != nil {
		return nil, err
	}
	info, err := parseVdiskInfo(*xmlresponse)
	if err != nil {
		return nil, err
	}
	for name, stat := range stats {
		stat.Health = info[name].Health
		stat.TotalDataTransferred = info[name].TotalDataTransferred
		stat.Status = info[name].Status
		stats[name] = stat
	}
	return stats, nil
}

func (c Client) GetSensorStatus() (map[string]SensorStatus, error) {
	xmlresponse, err := c.readXMLData("show/sensor-status")
	if err != nil {
		return nil, err
	}
	return parseSensorStatus(*xmlresponse)
}

func (c Client) readXMLData(command string) (*apiResponse, error) {
	var result apiResponse
	resp, err := c.httpclient.Get(c.server + "/api/" + command)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
