package p2000

import (
	"net/url"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/michep/snap-plugin-collector-p2000/client"
)

const (
	PluginName    = "p2000"
	PluginVersion = 1
	PluginVedor   = "mfms"
	param_server  = "server"
	param_authstr = "authstr"
)

type p2000Statistics interface {
	GetMetricNamespaces() []plugin.Namespace
	GetMetricValues(plugin.Metric, time.Time, *client.Client) ([]plugin.Metric, error)
	Reset()
}

type Plugin struct {
	client *client.Client
	system string
	stats  []p2000Statistics
}

func NewCollector(stats ...p2000Statistics) *Plugin {
	return &Plugin{stats: stats}
}

func (p *Plugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{PluginVedor, PluginName}, param_server, true)
	policy.AddNewStringRule([]string{PluginVedor, PluginName}, param_authstr, true)
	return *policy, nil
}

func (p *Plugin) GetMetricTypes(plugin.Config) ([]plugin.Metric, error) {
	var mts []plugin.Metric

	for _, stat := range p.stats {
		for _, namespace := range stat.GetMetricNamespaces() {
			mts = append(mts, plugin.Metric{Namespace: namespace})
		}
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
		p.client = client.NewClient(server, authstr)
	}

	now := time.Now()

	err := p.client.Login()
	if err != nil {
		return nil, err
	}

	for _, stat := range p.stats {
		stat.Reset()
	}

	for _, metric := range metrics {
		metric.Tags = map[string]string{"system": p.system}

		for _, stat := range p.stats {
			m, err := stat.GetMetricValues(metric, now, p.client)
			if err != nil {
				return nil, err
			}
			mts = append(mts, m...)
		}
	}

	return mts, nil
}
