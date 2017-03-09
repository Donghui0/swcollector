package funcs

import (
	"github.com/Donghui0/swcollector/g"
	"github.com/leancloud/satori/common/model"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	interval := g.Config().Transfer.Interval
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				SwIfMetrics,
				//CpuMetrics,
				//MemMetrics,
				//PingMetrics,
				//ConnMetrics,
			},
			Interval: interval,
		},
	}
}
