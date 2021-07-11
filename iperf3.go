package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type Iperf3 struct {
	interval int64          `toml:"interval"`
	flush_interval int64   `toml:"flush_interval"`
	flush_jitter    int64  `toml:"flush_jitter"`
metric_batch_size  int64   `toml:"url"`
	
}

const iperf3Config = `
## interval = "3600s"
## flush_interval = "1800s"
##  flush_jitter = "1800s"
##  metric_batch_size = 100
`


func (s* Iperf3) SampleConfig{} string{
	return  iperf3Config

func (s* Iperf3) Description SampleConfig() string {
	return " gather and return the metrics"
}

func Gather(s* Iperf3)(acc telegraf.Accumulator) error{
	
	return nil
}


func init(){
	inputs.Add("Iperf3", func telegraf.input {
		return & Iperf3{}
	})
}

