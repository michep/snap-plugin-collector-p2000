package p2000

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"time"
)

func (p Plugin) createControllerNamespaces() []plugin.Namespace {
	var ns []plugin.Namespace
	metrics := []string{"iops", "bytespersecond", "numberofreads", "numberofwrites", "dataread", "datawritten", "cpuload"}
	for _, m := range metrics {
		namespace := plugin.NewNamespace(PluginVedor, PluginName, "controller")
		namespace = namespace.AddDynamicElement("name", "component name")
		namespace = namespace.AddStaticElement(m)
		ns = append(ns, namespace)
	}
	return ns
}

func (p *Plugin) getControllerMetricValues(metric plugin.Metric, now time.Time) ([]plugin.Metric, error) {
	var err error
	var mts []plugin.Metric
	if p.ctlstat == nil {
		p.ctlstat, err = p.client.GetControllerStatistics()
		if err != nil {
			return nil, err
		}
	}

	for name, stat := range p.ctlstat {
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
		case "cpuload":
			m.Data = stat.CPULoad
		}
		mts = append(mts, m)
	}

	return mts, nil
}
