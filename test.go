package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type Test struct {
	CustomerName string `toml:"CustomerName"`
	UseCaseName string `toml:"UseCaseName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const testtype TestConfig = `
## CustomerName = "Latence"
## UseCaseName = "GatherData"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* testtype Test
) SampleConfig{} string{
	return  testtype Test
Config
}

func (s* testtype Test
) Description SampleConfig() string {
	return " Gather data and output it"
}

func Gather(s* testtype Test
)(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("testtype Test
", func telegraf.input {
		return & testtype Test{}
	})
}

