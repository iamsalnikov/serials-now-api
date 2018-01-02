package subscriptions

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
	var translatorID int64 = 2

	form := &url.Values{}
	form.Set("ID", strconv.FormatInt(serialID, 10))
	form.Set("T", strconv.FormatInt(translatorID, 10))

	expectedString := form.Encode()

	endpoint := NewEndpoint(serialID, translatorID)
	request, err := endpoint.BuildHttpRequest()
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	if string(requestBody) != expectedString {
		t.Errorf("I expected body \"%s\" but got \"%s\"", expectedString, string(requestBody))
	}
}

func TestEndpoint_ParseResponse(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Errorf("Couldn't create client: %s", err)
	}

	endpoint := NewEndpoint(10, 2)
	err = endpoint.ParseResponse(response)

	if err != nil {
		t.Errorf("I got undexpected error: %s", err)
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

	endpoint := NewEndpoint(10, 2)
	err = endpoint.ParseResponse(response)

	if err != UnexpectedStatusCode {
		t.Errorf("I expected error \"%s\" but got \"%s\"", UnexpectedStatusCode, err)
	}
}

func TestEndpoint_ParseResponseBadRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Response":"Неверный запрос."}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Errorf("Couldn't create client: %s", err)
	}

	endpoint := NewEndpoint(10, 2)
	err = endpoint.ParseResponse(response)

	if err != BadRequest {
		t.Errorf("I expected error \"%s\" but got \"%s\"", UnexpectedStatusCode, err)
	}
}
