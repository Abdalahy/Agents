package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type ORE struct {
	CustomerName string `toml:"CustomerName"`
	SiteName string `toml:"	SiteName "`
	UseCaseName string `toml:"UseCaseName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const OREConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
## UseCaseName = "GatherData"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* ORE) SampleConfig{} string{
	return  OREConfig
}

func (s* ORE) Description SampleConfig() string {
	return " Gather weather data and output it"
}

func Gather(s* ORE)(acc telegraf.Accumulator) error{
	
	return nil
}


func init(){
	inputs.Add("ORE", func telegraf.input {
		return & ORE{}
	})
}

