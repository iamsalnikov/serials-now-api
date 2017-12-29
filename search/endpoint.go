package search

import (
	"net/http"
	"fmt"
)

type Endpoint struct {}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (r *Endpoint) BuildHttpRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/json/Search-2.json", nil)
}

func (r *Endpoint) ParseResponse(response *http.Response) error {
	fmt.Println("Parse response")
	return nil
}