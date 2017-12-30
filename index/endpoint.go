package index

import (
	"encoding/json"
	"errors"
	"net/http"
)

var UnexpectedStatusCode = errors.New("unexpected status code")

type Endpoint struct {
	Serials [][]Serial
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/json/Main-2.json", nil)
}

func (e *Endpoint) ParseResponse(response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return UnexpectedStatusCode
	}

	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(&e.Serials)
}
