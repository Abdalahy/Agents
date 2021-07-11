package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type LRE struct {
	CustomerName string `toml:"CustomerName"`
	SiteName string `toml:"	SiteName "`
	UseCaseName string `toml:"UseCaseName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const LREConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
## UseCaseName = "GatherData"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* LRE) SampleConfig{} string{
	return  LREConfig
}

func (s* LRE) Description SampleConfig() string {
	return " Gather latency metric data and output it"
}

func Gather(s* LRE)(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("LRE", func telegraf.input {
		return & LRE
	})
}

