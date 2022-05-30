package workers

import (
	"datapipe/sinks"
	"datapipe/sources"
)

func Transfer(src sources.Source, sink sinks.Sink) error {
	work := src.List()

	for _, item := range work {
		sink.Store(item)
		src.Update(item, "transfered", "true")
	}
}
