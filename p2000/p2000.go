package p2000

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/michep/snap-plugin-collector-p2000/parser"
	"net/url"
	"time"
)

const (
	PluginName    = "p2000"
	PluginVersion = 1
	PluginVedor   = "mfms"
	param_server  = "server"
	param_authstr = "authstr"
)

type Plugin struct {
	client       *parser.Client
	system       string
	diskstat     map[string]parser.DiskStatistics
	vdiskstat    map[string]parser.VdiskStatistics
	ctlstat      map[string]parser.ControllerStatistics
	sensorstat   map[string]parser.SensorStatus
	hostportstat map[string]parser.HostPortStatistics
}

func NewCollector() *Plugin {
	return &Plugin{}
}

func (p *Plugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{PluginVedor, PluginName}, param_server, true)
	policy.AddNewStringRule([]string{PluginVedor, PluginName}, param_authstr, true)
	return *policy, nil
}

func (p *Plugin) GetMetricTypes(plugin.Config) ([]plugin.Metric, error) {
	var mts []plugin.Metric

	ns := p.createDiskNamespaces()
	for _, namespace := range ns {
		mts = append(mts, plugin.Metric{Namespace: namespace})
	}
	ns = p.createVdiskNamespaces()
	for _, namespace := range ns {
		mts = append(mts, plugin.Metric{Namespace: namespace})
	}
	ns = p.createControllerNamespaces()
	for _, namespace := range ns {
		mts = append(mts, plugin.Metric{Namespace: namespace})
	}
	ns = p.createSensorStatusNamespaces()
	for _, namespace := range ns {
		mts = append(mts, plugin.Metric{Namespace: namespace})
	}
	ns = p.createHostPortNamespaces()
	for _, namespace := range ns {
		mts = append(mts, plugin.Metric{Namespace: namespace})
	}
	return mts, nil
}

func (p *Plugin) CollectMetrics(metrics []plugin.Metric) ([]plugin.Metric, error) {
	var mts []plugin.Metric

	if p.client == nil {
		server, _ := metrics[0].Config.GetString(param_server)
		authstr, _ := metrics[0].Config.GetString(param_authstr)
		u, err := url.Parse(server)
		if err != nil {
			return nil, err
		}
		p.system = u.Hostname()
		p.client = parser.NewClient(server, authstr)
	}

	p.diskstat = nil
	p.vdiskstat = nil
	p.ctlstat = nil
	p.sensorstat = nil
	p.hostportstat = nil

	now := time.Now()

	err := p.client.Login()
	if err != nil {
		return nil, err
	}

	for _, metric := range metrics {
		metric.Tags = map[string]string{"system": p.system}
		switch metric.Namespace[2].Value {
		case "drive":
			m, err := p.getDiskMetricValues(metric, now)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		case "vdisk":
			m, err := p.getVdiskMetricValues(metric, now)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		case "controller":
			m, err := p.getControllerMetricValues(metric, now)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		case "sensor":
			m, err := p.getSensorStatusMetricValues(metric, now)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		case "hostport":
			m, err := p.getHostPortkMetricValues(metric, now)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		}
	}

	return mts, nil
}
