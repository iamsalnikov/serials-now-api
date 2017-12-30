package test

import (
	"errors"
	"net/http"
)

var StubEndpointError = errors.New("stub endpoint error")

type StubEndpoint struct {
	ErrorOnRequest       bool
	ErrorOnParseResponse bool
}

func (e *StubEndpoint) BuildHttpRequest() (*http.Request, error) {
	request, _ := http.NewRequest(http.MethodGet, "/stub", nil)

	err := StubEndpointError
	if !e.ErrorOnRequest {
		err = nil
	}

	return request, err
}

func (e *StubEndpoint) ParseResponse(response *http.Response) error {
	if e.ErrorOnParseResponse {
		return StubEndpointError
	}

	return nil
}
