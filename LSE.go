package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type LSE struct {
	CustomerName string `toml:"CustomerName"`
	SiteName string `toml:"	SiteName "`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
}

const LSEConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* LSE ) SampleConfig{} string{
	return  LSConfig
}

func (s* LSE
) Description SampleConfig() string {
	return " Gather latency metric data and output it"
}

func Gather(s* LSE )(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("LSE", func telegraf.input {
		return & LSE{}
	
	})
}

