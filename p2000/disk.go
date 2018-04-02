package p2000

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"time"
)

func (p Plugin) createDiskNamespaces() []plugin.Namespace {
	var ns []plugin.Namespace
	metrics := []string{"iops", "bytespersecond", "numberofreads", "numberofwrites", "dataread", "datawritten", "totaldatatransferred", "health", "queuedepth"}
	for _, m := range metrics {
		namespace := plugin.NewNamespace(PluginVedor, PluginName, "drive")
		namespace = namespace.AddDynamicElement("name", "component name")
		namespace = namespace.AddStaticElement(m)
		ns = append(ns, namespace)
	}
	return ns
}

func (p *Plugin) getDiskMetricValues(metric plugin.Metric, now time.Time) ([]plugin.Metric, error) {
	var err error
	var mts []plugin.Metric
	if p.diskstat == nil {
		p.diskstat, err = p.client.GetDiskStatistics()
		if err != nil {
			return nil, err
		}
	}

	for name, stat := range p.diskstat {
		ns := plugin.NewNamespace()
		for _, nse := range metric.Namespace {
			ns = append(ns, nse)
		}
		m := plugin.Metric{Namespace: ns, Timestamp: now}
		m.Namespace[3].Value = name
		switch m.Namespace[4].Value {
		case "iops":
			m.Data = stat.Iops
		case "bytespersecond":
			m.Data = stat.BytesPerSecond
		case "numberofreads":
			m.Data = stat.NumberOfReads
		case "numberofwrites":
			m.Data = stat.NumberOfWrites
		case "dataread":
			m.Data = stat.DataRead
		case "datawritten":
			m.Data = stat.DataWritten
		case "totaldatatransferred":
			m.Data = stat.TotalDataTransferred
		case "health":
			m.Data = stat.Health
		case "queuedepth":
			m.Data = stat.QueueDepth
		}
		mts = append(mts, m)
	}

	return mts, nil
}
