package main

import (
	"datapipe/sinks"
	"datapipe/sources"
	"datapipe/workers"
	"os"

	"github.com/segmentio/conf"
)

func main() {
	var config struct {
		SourceTable string `conf:"source_table`
	}
	conf.Load(&config)

	airtable := sources.NewAirTableClient(os.Getenv("airtable"), os.Getenv("baseid"), config.SourceTable)
	planetScale := sinks.NewPlanetScaleClient(os.Getenv("DSN"))
	workers.Transfer(airtable, planetScale)
}
