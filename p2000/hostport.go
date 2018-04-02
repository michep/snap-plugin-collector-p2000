package p2000

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

func (p Plugin) createHostPortNamespaces() []plugin.Namespace {
	ns := []plugin.Namespace{}
	metrics := []string{"iops", "bytespersecond", "numberofreads", "numberofwrites", "dataread", "datawritten", "queuedepth"}
	for _, m := range metrics {
		namespace := plugin.NewNamespace(PluginVedor, PluginName, "hostport")
		namespace = namespace.AddDynamicElement("name", "component name")
		namespace = namespace.AddStaticElement(m)
		ns = append(ns, namespace)
	}
	return ns
}
