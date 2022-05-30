package sources

import "datapipe/data"

type Source interface {
	List() []data.Record
	Get()
	Update(record data.Record, attribute, value string)
}
