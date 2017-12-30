package search

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoint_ParseResponse(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"path":           "/not-empty-search.json",
			"expectedLength": 2,
		},
		{
			"path":           "/empty-search.json",
			"expectedLength": 0,
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		filePath := "testdata" + r.URL.Path

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	for _, testCase := range testCases {
		path := testCase["path"].(string)
		expectedLength := testCase["expectedLength"].(int)

		response, err := http.Get(server.URL + path)
		if err != nil {
			t.Errorf("Couldn't sent request: %s", err)
		}

		endpoint := NewEndpoint()
		err = endpoint.ParseResponse(response)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(endpoint.Serials) != expectedLength {
			t.Errorf("I got unexpected serials length. Expected length: %d, got: %d", expectedLength, len(endpoint.Serials))
		}
	}
}

func TestEndpoint_ParseResponseBadStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	response, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Couldn't send request: %s", err)
	}

	endpoint := NewEndpoint()
	err = endpoint.ParseResponse(response)

	if err != UnexpectedStatusCode {
		t.Errorf("I expected error UnexpectedStatusCode but got \"%s\"", err)
	}
}
