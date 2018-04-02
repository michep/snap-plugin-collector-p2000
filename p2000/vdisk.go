package p2000

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"time"
)

func (p Plugin) createVdiskNamespaces() []plugin.Namespace {
	var ns []plugin.Namespace
	metrics := []string{"iops", "bytespersecond", "numberofreads", "numberofwrites", "dataread", "datawritten", "totaldatatransferred", "health", "status", "avgrsptime", "avgreadrsptime", "avgwritersptime"}
	for _, m := range metrics {
		namespace := plugin.NewNamespace(PluginVedor, PluginName, "vdisk")
		namespace = namespace.AddDynamicElement("name", "component name")
		namespace = namespace.AddStaticElement(m)
		ns = append(ns, namespace)
	}
	return ns
}

func (p *Plugin) getVdiskMetricValues(metric plugin.Metric, now time.Time) ([]plugin.Metric, error) {
	var err error
	var mts []plugin.Metric
	if p.vdiskstat == nil {
		p.vdiskstat, err = p.client.GetVdiskStatistics()
		if err != nil {
			return nil, err
		}
	}

	for name, stat := range p.vdiskstat {
		ns := plugin.NewNamespace()
		tags := make(map[string]string)
		ns = append(ns, metric.Namespace...)
		for k, v := range metric.Tags {
			tags[k] = v
		}
		m := plugin.Metric{Namespace: ns, Timestamp: now, Tags: tags}
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
		case "status":
			m.Data = stat.Status
		case "avgrsptime":
			m.Data = stat.AvgRspTime
		case "avgreadrsptime":
			m.Data = stat.AvgReadRspTime
		case "avgwritersptime":
			m.Data = stat.AvgWriteRspTime
		}
		mts = append(mts, m)
	}

	return mts, nil
}
