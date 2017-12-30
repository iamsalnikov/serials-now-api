package index

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoint_ParseResponse(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"path":            "/not-empty.json",
			"expectedLengths": []int{2, 0, 1, 2},
		},
		{
			"path":            "/empty.json",
			"expectedLengths": []int{0, 0, 0, 0},
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
		expectedLengths := testCase["expectedLengths"].([]int)

		response, err := http.Get(server.URL + path)
		if err != nil {
			t.Errorf("Couldn't sent request: %s", err)
		}

		endpoint := NewEndpoint()
		err = endpoint.ParseResponse(response)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		for i, expectedLength := range expectedLengths {
			if len(endpoint.Serials[i]) != expectedLength {
				t.Errorf("I got unexpected serials length. Expected length: %d, got: %d", expectedLength, len(endpoint.Serials[i]))
			}
		}
	}
}

func TestEndpoint_ParseResponseBadStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	response, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Couldn't send request (%s)", err)
	}

	endpoint := NewEndpoint()
	err = endpoint.ParseResponse(response)
	if err != UnexpectedStatusCode {
		t.Errorf("Expected error UnexpectedStatusCode, got: \"%s\"", err)
	}
}
