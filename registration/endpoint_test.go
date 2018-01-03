package registration

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestEndpoint_BuildHttpRequest(t *testing.T) {
	form := &url.Values{}
	form.Set("Email", "hello@world")
	form.Set("Login", "hellologin")
	form.Set("Password", "hellopassword")

	endpoint := NewEndpoint(form.Get("Email"), form.Get("Login"), form.Get("Password"))
	request, err := endpoint.BuildHttpRequest()
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}

	expectedString := form.Encode()
	if string(body) != expectedString {
		t.Fatalf("I expected body \"%s\" but got \"%s\"", expectedString, string(body))
	}
}

func TestEndpoint_ParseResponseBadStatusCode(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}

	endpoint := NewEndpoint("", "", "")
	err = endpoint.ParseResponse(response)
	if err != UnexpectedStatusCode {
		t.Fatalf("I expected error \"%s\" but got \"%s\"", UnexpectedStatusCode, err)
	}
}

func TestEndpoint_ParseResponseBadResponse(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Response": "Hello world"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}

	endpoint := NewEndpoint("", "", "")
	err = endpoint.ParseResponse(response)
	if err != BadRequest {
		t.Fatalf("I expected error \"%s\" but got \"%s\"", BadRequest, err)
	}
}

func TestEndpoint_ParseResponse(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	response, err := http.Post(server.URL, "test", nil)
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}

	endpoint := NewEndpoint("", "", "")
	err = endpoint.ParseResponse(response)
	if err != nil {
		t.Fatalf("I got unexpected error: %s", err)
	}
}
