package sources

import (
	"datapipe/data"
	"fmt"
	"net/http"
	"time"
)

const majorAPIVersion = 0
const retryDelayIfRateLimited = 5 * time.Second

var apiBaseURL = fmt.Sprintf("https://api.airtable.com/v%d", majorAPIVersion)

type AirTableClient struct {
	apiKey                   string
	baseID                   string
	table                    string
	ShouldRetryIfRateLimited bool
	HTTPClient               *http.Client
}

func NewAirTableClient(apiKey, baseID, table string) *AirTableClient {
	c := AirTableClient{
		apiKey:                   apiKey,
		baseID:                   baseID,
		table:                    table,
		ShouldRetryIfRateLimited: true,
		HTTPClient:               http.DefaultClient,
	}
	return &c
}

func (c *AirTableClient) List() []data.Record {
	panic("implement me...")
}

func (c *AirTableClient) Get() {
	panic("implement me...")
}

func (c *AirTableClient) Update(record data.Record, attribute, value string) {
	panic("implement me...")
}
