package user

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoint_ParseResponseBadStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	response, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	endpoint := NewEndpoint()
	err = endpoint.ParseResponse(response)
	if err != UnexpectedStatusCode {
		t.Errorf("I expected error \"%s\" but got \"%s\"", UnexpectedStatusCode, err)
	}
}

func TestEndpoint_ParseResponseLongAndShortAnswers(t *testing.T) {
	testCases := map[string][]byte{
		"/zero":  []byte(`[]`),
		"/one":   []byte(`[[]]`),
		"/five":  []byte(`[[], [], [], [], []]`),
		"/seven": []byte(`[[], [], [], [], [], [], []]`),
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(testCases[r.URL.Path])
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	for path := range testCases {
		response, err := http.Get(server.URL + path)
		if err != nil {
			t.Errorf("I got unexpected error: %s", err)
		}

		endpoint := NewEndpoint()
		err = endpoint.ParseResponse(response)
		if err != UnexpectedAnswerLength {
			t.Errorf("I expected error \"%s\" but got \"%s\"", UnexpectedAnswerLength, err)
		}
	}
}

func TestEndpoint_ParseResponse(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"path":       "/empty.json",
			"favorites":  0,
			"watched":    0,
			"new":        0,
			"subscribed": 0,
			"voted":      0,
			"comments":   0,
		},
		{
			"path":       "/each-by-one.json",
			"favorites":  1,
			"watched":    1,
			"new":        1,
			"subscribed": 1,
			"voted":      1,
			"comments":   1,
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
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
		response, err := http.Get(server.URL + testCase["path"].(string))
		if err != nil {
			t.Errorf("I got unexpected error: %s", err)
		}

		endpoint := NewEndpoint()
		err = endpoint.ParseResponse(response)
		if err != nil {
			t.Errorf("I got unexpected error: %s", err)
		}

		expectedCount := testCase["favorites"].(int)
		actualCount := len(endpoint.Data.FavoriteEpisodes)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}

		expectedCount = testCase["watched"].(int)
		actualCount = len(endpoint.Data.WatchedEpisodes)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}

		expectedCount = testCase["new"].(int)
		actualCount = len(endpoint.Data.NewEpisodes)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}

		expectedCount = testCase["subscribed"].(int)
		actualCount = len(endpoint.Data.SubscribedSerials)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}

		expectedCount = testCase["voted"].(int)
		actualCount = len(endpoint.Data.VotedSerials)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}

		expectedCount = testCase["comments"].(int)
		actualCount = len(endpoint.Data.Comments)
		if actualCount != expectedCount {
			t.Errorf("I expected count %d but got %d", expectedCount, actualCount)
		}
	}
}
