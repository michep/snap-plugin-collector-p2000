package p2000

import (
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/michep/snap-plugin-collector-p2000/client"
)

type VdiskStatistics struct {
	stats map[string]client.VdiskStatistics
}

func (s VdiskStatistics) GetMetricNamespaces() []plugin.Namespace {
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

func (s *VdiskStatistics) GetMetricValues(metric plugin.Metric, now time.Time, client *client.Client) ([]plugin.Metric, error) {
	var err error
	var mts []plugin.Metric

	if metric.Namespace[2].Value != "vdisk" {
		return nil, nil
	}

	if s.stats == nil {
		s.stats, err = client.GetVdiskStatistics()
		if err != nil {
			return nil, err
		}
	}

	for name, stat := range s.stats {
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

func (s *VdiskStatistics) Reset() {
	s.stats = nil
}
