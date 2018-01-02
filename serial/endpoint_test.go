package serial

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestEndpoint_BuildHttpRequest(t *testing.T) {
	var serialID int64 = 10
	endpoint := NewEndpoint(serialID)

	req, err := endpoint.BuildHttpRequest()
	if err != nil {
		t.Errorf("Couldn't create request: %s", err)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Couldn't read body of request: %s", err)
	}

	form := &url.Values{}
	form.Set("ID", strconv.FormatInt(serialID, 10))

	if form.Encode() != string(data) {
		t.Errorf("I expected request data %s but got %s", form.Encode(), string(data))
	}
}

func TestEndpoint_ParseResponseBadStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Errorf("Couldn't create client: %s", err)
	}

	endpoint := NewEndpoint(10)
	err = endpoint.ParseResponse(response)

	if err != UnexpectedStatusCode {
		t.Errorf("I expected error \"%s\" but got \"%s\"", UnexpectedStatusCode, err)
	}
}

func TestEndpoint_ParseResponse(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"path":                     "/auth-answer.json",
			"watchHistory":             1,
			"translationSubscriptions": 1,
		},
		{
			"path":                     "/unauth-answer.json",
			"watchHistory":             0,
			"translationSubscriptions": 0,
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
		endpoint := NewEndpoint(0)

		response, err := http.Post(server.URL+testCase["path"].(string), "test", nil)
		if err != nil {
			t.Errorf("I got unexpected error: %s", err)
		}

		err = endpoint.ParseResponse(response)
		if err != nil {
			t.Errorf("I got unexpected error: %s", err)
		}

		expectedLength := testCase["watchHistory"].(int)
		actualLength := len(endpoint.Data.WatchHistory)
		if expectedLength != actualLength {
			t.Errorf("I expected length %d but got %d", expectedLength, actualLength)
		}

		expectedLength = testCase["translationSubscriptions"].(int)
		actualLength = len(endpoint.Data.TranslationSubscriptions)
		if expectedLength != actualLength {
			t.Errorf("I expected length %d but got %d", expectedLength, actualLength)
		}
	}
}
