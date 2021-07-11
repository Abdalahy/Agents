package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type LatencyRealSoultion struct {
	CustomerName string `toml:"CustomerName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const LatencyRealSoultionConfig = `
## CustomerName = "Latence"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* LatencyRealSoultion) SampleConfig{} string{
	return  LatencyRealSoultionConfig
}

func (s* LatencyRealSoultion) Description SampleConfig() string {
	return " Gather data and output it"
}

func Gather(s* LatencyRealSoultion)(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("LatencyRealSoultion", func telegraf.input {
		return & LatencyRealSoultion{}
	})
}

