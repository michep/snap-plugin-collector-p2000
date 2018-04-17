package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/michep/snap-plugin-collector-p2000/p2000"
)

func main() {
	coll := p2000.NewCollector(
		&p2000.DiskStatistics{},
		&p2000.VdiskStatistics{},
		&p2000.ControllerStatistics{},
		&p2000.SensorStatus{},
		&p2000.HostportStatistics{},
	)
	plugin.StartCollector(coll, p2000.PluginName, p2000.PluginVersion, plugin.RoutingStrategy(plugin.StickyRouter))
}
