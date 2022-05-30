package sinks

import "datapipe/data"

type Sink interface {
	Store(record data.Record)
}
