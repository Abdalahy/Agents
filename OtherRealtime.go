package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type OtherRealtime struct {
	CustomerName string `toml:"CustomerName"`
	SiteName string `toml:"	SiteName "`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
}

const OtherRealtimeConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* OtherRealtime ) SampleConfig{} string{
	return  LSConfig
}

func (s* OtherRealtime
) Description SampleConfig() string {
	return " Gather latency metric data and output it"
}

func Gather(s* OtherRealtime )(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("OtherRealtime", func telegraf.input {
		return & OtherRealtime{}
	
	})
}

