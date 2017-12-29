package search

import (
	"net/http"
	"encoding/json"
)

type Endpoint struct {
	Serials []Serial
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/json/Search-2.json", nil)
}

func (e *Endpoint) ParseResponse(response *http.Response) error {
	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(&e.Serials)
}

