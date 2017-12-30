package serials_now_api

import (
	"github.com/iamsalnikov/serials-now-api/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SendWithBadBuiltRequest(t *testing.T) {
	endpoint := &test.StubEndpoint{
		ErrorOnRequest: true,
	}

	client, err := NewClient("http://example.com")
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	err = client.Send(endpoint)
	if err != test.StubEndpointError {
		t.Errorf("I expected StubEndpointError but got \"%s\"", err)
	}
}

func TestClient_SendWithBadParseResponse(t *testing.T) {
	endpoint := &test.StubEndpoint{
		ErrorOnParseResponse: true,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	client, err := NewClient(server.URL)
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	err = client.Send(endpoint)
	if err != test.StubEndpointError {
		t.Errorf("I expected StubEndpointError but got \"%s\"", err)
	}
}
