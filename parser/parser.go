package parser

import (
	"strconv"
)

//import "encoding/xml"

type apiResponse struct {
	//XMLName xml.Name `xml:"RESPONSE"`
	Objects []apiResponseObject `xml:"OBJECT"`
}

type apiResponseObject struct {
	//XMLName xml.Name `xml:"OBJECT"`
	Basetype   string                      `xml:"basetype,attr"`
	Properties []apiResponseObjectProperty `xml:"PROPERTY"`
}

type apiResponseObjectProperty struct {
	//XMLName xml.Name `xml:"PROPERTY"`
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type statistics struct {
	Iops           int64
	BytesPerSecond int64
	NumberOfReads  int64
	NumberOfWrites int64
	DataRead       int64
	DataWritten    int64
}

type DiskStatistics struct {
	Name         string
	SerialNumber string
	QueueDepth   int64
	statistics
	diskInfo
}

type VdiskStatistics struct {
	Name            string
	SerialNumber    string
	AvgRspTime      int64
	AvgReadRspTime  int64
	AvgWriteRspTime int64
	statistics
	vdiskInfo
}

type ControllerStatistics struct {
	Name    string
	CPULoad int64
	statistics
}

type HostPortStatistics struct {
	Name       string
	QueueDepth int64
	statistics
}

type diskInfo struct {
	Location             string
	TotalDataTransferred int64
	Health               int64
}

type vdiskInfo struct {
	Status               int64
	TotalDataTransferred int64
	Health               int64
}

type hostPortInfo struct {
	Health int64
}

type StatusInfo struct {
	SessionKey string
}

func parseControllerStatistics(resp apiResponse) (map[string]ControllerStatistics, error) {
	res := make(map[string]ControllerStatistics)
	for _, obj := range resp.Objects {
		if obj.Basetype == "controller-statistics" {
			stat := ControllerStatistics{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "durable-id":
					stat.Name = prop.Value
				case "cpu-load":
					stat.CPULoad, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "iops":
					stat.Iops, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "bytes-per-second-numeric":
					stat.BytesPerSecond, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-reads":
					stat.NumberOfReads, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-writes":
					stat.NumberOfWrites, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-read-numeric":
					stat.DataRead, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-written-numeric":
					stat.DataWritten, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[stat.Name] = stat
		}
	}
	return res, nil
}

func parseHostPortStatistics(resp apiResponse) (map[string]HostPortStatistics, error) {
	res := make(map[string]HostPortStatistics)
	for _, obj := range resp.Objects {
		if obj.Basetype == "host-port-statistics" {
			stat := HostPortStatistics{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "durable-id":
					stat.Name = prop.Value
				case "queue-depth":
					stat.QueueDepth, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "iops":
					stat.Iops, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "bytes-per-second-numeric":
					stat.BytesPerSecond, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-reads":
					stat.NumberOfReads, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-writes":
					stat.NumberOfWrites, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-read-numeric":
					stat.DataRead, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-written-numeric":
					stat.DataWritten, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[stat.Name] = stat
		}
	}
	return res, nil
}

func parseDiskStatistics(resp apiResponse) (map[string]DiskStatistics, error) {
	res := make(map[string]DiskStatistics)
	for _, obj := range resp.Objects {
		if obj.Basetype == "disk-statistics" {
			stat := DiskStatistics{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "durable-id":
					stat.Name = prop.Value
				case "serial-number":
					stat.SerialNumber = prop.Value
				case "queue-depth":
					stat.QueueDepth, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "iops":
					stat.Iops, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "bytes-per-second-numeric":
					stat.BytesPerSecond, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-reads":
					stat.NumberOfReads, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-writes":
					stat.NumberOfWrites, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-read-numeric":
					stat.DataRead, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-written-numeric":
					stat.DataWritten, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[stat.Name] = stat
		}
	}
	return res, nil
}

func parseDiskInfo(resp apiResponse) (map[string]diskInfo, error) {
	var name string
	res := make(map[string]diskInfo)
	for _, obj := range resp.Objects {
		if obj.Basetype == "drives" {
			info := diskInfo{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "durable-id":
					name = prop.Value
				case "location":
					info.Location = prop.Value
				case "total-data-transferred-numeric":
					info.TotalDataTransferred, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "health-numeric":
					info.Health, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[name] = info
		}
	}
	return res, nil
}

func parseVdiskStatistics(resp apiResponse) (map[string]VdiskStatistics, error) {
	res := make(map[string]VdiskStatistics)
	for _, obj := range resp.Objects {
		if obj.Basetype == "vdisk-statistics" {
			stat := VdiskStatistics{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "name":
					stat.Name = prop.Value
				case "serial-number":
					stat.SerialNumber = prop.Value
				case "iops":
					stat.Iops, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "bytes-per-second-numeric":
					stat.BytesPerSecond, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-reads":
					stat.NumberOfReads, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "number-of-writes":
					stat.NumberOfWrites, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-read-numeric":
					stat.DataRead, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "data-written-numeric":
					stat.DataWritten, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[stat.Name] = stat
		}
	}
	return res, nil
}

func parseVdiskInfo(resp apiResponse) (map[string]vdiskInfo, error) {
	var name string
	res := make(map[string]vdiskInfo)
	for _, obj := range resp.Objects {
		if obj.Basetype == "virtual-disks" {
			info := vdiskInfo{}
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "name":
					name = prop.Value
				case "status-numeric":
					info.Status, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "total-data-transferred-numeric":
					info.TotalDataTransferred, _ = strconv.ParseInt(prop.Value, 10, 64)
				case "health-numeric":
					info.Health, _ = strconv.ParseInt(prop.Value, 10, 64)
				}
			}
			res[name] = info
		}
	}
	return res, nil
}

func parseStatusInfo(resp apiResponse) (StatusInfo, error) {
	var res StatusInfo
	for _, obj := range resp.Objects {
		if obj.Basetype == "status" {
			for _, prop := range obj.Properties {
				switch prop.Name {
				case "response":
					res.SessionKey = prop.Value
				}
			}
		}
	}
	return res, nil
}
