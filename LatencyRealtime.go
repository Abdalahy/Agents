package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type LatencyRealtime struct {
	CustomerName string `toml:"CustomerName"`
	UseCaseName string `toml:"UseCaseName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const LatencyRealtimeConfig = `
## CustomerName = "Latence"
## UseCaseName = "GatherData"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* LatencyRealtime) SampleConfig{} string{
	return  LatencyRealtimeConfig
}

func (s* LatencyRealtime) Description SampleConfig() string {
	return " Gather data and output it"
}

func Gather(s* LatencyRealtime)(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("LatencyRealtime", func telegraf.input {
		return & LatencyRealtime{}
	})
}

