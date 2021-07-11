package main

	import (
	// "github.com/influxdata/telegraf"
	// "github.com/influxdata/telegraf/plugins/inputs"	
)

type Env struct {
	CustomerName string `toml:"CustomerName"`
	SiteName string `toml:"	SiteName "`
	UseCaseName string `toml:"UseCaseName"`
	LatenceTechInfo string `toml:"LatenceTechInfo"`
	
}

const envConfig = `
## CustomerName = "Latence"
## SiteName = "Sitename"
## UseCaseName = "GatherData"
##  LatenceTechInfo = "www.Latence.co"

`


func (s* Env) SampleConfig{} string{
	return  envConfig
}

func (s* Env) Description SampleConfig() string {
	return " Gather weather data and output it"
}

func Gather(s* Env)(acc telegraf.Accumulator) error{
	

	return nil
}


func init(){
	inputs.Add("Env", func telegraf.input {
		return & Env{}
	})
}

