package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type PingSize struct {
	
}

const PingSizeConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
##  LatenceTechInfo = "www.Latence.co"

`

func (s* PingSize ) SampleConfig{} string{
	return  LSConfig
}

func (s* PingSize
) Description SampleConfig() string {
	return " Gather latency metric data and output it"
}

func Gather(s* PingSize )(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("PingSize", func telegraf.input {
		return & PingSize{}
	
	})
}

