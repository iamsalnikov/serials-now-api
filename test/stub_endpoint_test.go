package test

import (
	"net/http"
	"testing"
)

func TestStubEndpoint_BuildHttpRequestWithError(t *testing.T) {
	endpoint := &StubEndpoint{
		ErrorOnRequest: true,
	}

	_, err := endpoint.BuildHttpRequest()
	if err != StubEndpointError {
		t.Errorf("I expected StubEndpointError but got \"%s\"", err)
	}
}

func TestStubEndpoint_BuildHttpRequestWithoutError(t *testing.T) {
	endpoint := &StubEndpoint{
		ErrorOnRequest: false,
	}

	_, err := endpoint.BuildHttpRequest()
	if err != nil {
		t.Errorf("I expected no errors but got \"%s\"", err)
	}
}

func TestStubEndpoint_ParseResponseWithError(t *testing.T) {
	endpoint := &StubEndpoint{
		ErrorOnParseResponse: true,
	}

	response := &http.Response{}

	err := endpoint.ParseResponse(response)
	if err != StubEndpointError {
		t.Errorf("I expected StubEndpointError but got \"%s\"", err)
	}
}

func TestStubEndpoint_ParseResponseWithoutError(t *testing.T) {
	endpoint := &StubEndpoint{
		ErrorOnParseResponse: false,
	}

	response := &http.Response{}

	err := endpoint.ParseResponse(response)
	if err != nil {
		t.Errorf("I expected no errors but got \"%s\"", err)
	}
}
