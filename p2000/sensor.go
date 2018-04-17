package p2000

import (
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/michep/snap-plugin-collector-p2000/parser"
)

type SensorStatus struct {
	stats map[string]parser.SensorStatus
}

func (s SensorStatus) GetMetricNamespaces() []plugin.Namespace {
	var ns []plugin.Namespace
	metrics := []string{"status", "reading"}
	for _, m := range metrics {
		namespace := plugin.NewNamespace(PluginVedor, PluginName, "sensor")
		namespace = namespace.AddDynamicElement("name", "component name")
		namespace = namespace.AddDynamicElement("type", "component type")
		namespace = namespace.AddStaticElement(m)
		ns = append(ns, namespace)
	}
	return ns
}

func (s *SensorStatus) GetMetricValues(metric plugin.Metric, now time.Time, client *parser.Client) ([]plugin.Metric, error) {
	var err error
	var mts []plugin.Metric

	if metric.Namespace[2].Value != "sensor" {
		return nil, nil
	}

	if s.stats == nil {
		s.stats, err = client.GetSensorStatus()
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
		m.Namespace[4].Value = stat.Type
		switch m.Namespace[5].Value {
		case "status":
			m.Data = stat.Status
		case "reading":
			m.Data = stat.Value
		}
		mts = append(mts, m)
	}

	return mts, nil
}

func (s *SensorStatus) Reset() {
	s.stats = nil
}
