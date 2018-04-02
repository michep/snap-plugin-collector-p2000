package p2000

import (
	"fmt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"testing"
)

func TestPlugin_GetMetricTypes(t *testing.T) {
	p := NewCollector()
	cfg := plugin.NewConfig()
	mts, _ := p.GetMetricTypes(cfg)
	fmt.Printf("%+v\n", len(mts))
}

func TestPlugin_CollectMetrics(t *testing.T) {
	p := NewCollector()
	cfg := plugin.NewConfig()
	cfg[param_server] = "http://172.16.18.31:80"
	cfg[param_authstr] = "0e4997806bb599dec1864e034f9e59f9"
	mts, _ := p.GetMetricTypes(cfg)
	mts[0].Config = cfg

	metrics, _ := p.CollectMetrics(mts)
	//fmt.Printf("%+v\n", len(metrics))
	for _, m := range metrics {
		fmt.Printf("%+v\n", m)
	}
}
