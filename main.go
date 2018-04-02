package main

import (
	"github.com/michep/snap-plugin-collector-p2000/p2000"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func main() {
	plugin.StartCollector(p2000.NewCollector(), p2000.PluginName, p2000.PluginVersion, plugin.RoutingStrategy(plugin.StickyRouter))
}
