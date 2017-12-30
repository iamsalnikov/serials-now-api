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

func TestClient_SetCookieNew(t *testing.T) {
	client, err := NewClient("http://example.com")
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	client.SetCookie("test", "value")

	has := false
	for _, cookie := range client.client.Jar.Cookies(client.baseUri) {
		has = cookie.Name == "test" && cookie.Value == "value"
		if has {
			break
		}
	}

	if !has {
		t.Errorf("Can not find necessary cookie")
	}
}

func TestClient_SetCookieExists(t *testing.T) {
	client, err := NewClient("http://example.com")
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	client.SetCookie("test", "value-1")
	client.SetCookie("test", "value-2")

	has := false
	for _, cookie := range client.client.Jar.Cookies(client.baseUri) {
		has = cookie.Name == "test" && cookie.Value == "value-2"
		if has {
			break
		}
	}

	if !has {
		t.Errorf("Can not find necessary cookie")
	}
}

func TestClient_SetLogin(t *testing.T) {
	client, err := NewClient("http://example.com")
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	client.SetLogin("hello")

	has := false
	for _, cookie := range client.client.Jar.Cookies(client.baseUri) {
		has = cookie.Name == LOGIN_COOKIE_NAME && cookie.Value == "hello"
		if has {
			break
		}
	}

	if !has {
		t.Errorf("Can not find necessary cookie")
	}
}

func TestClient_SetPassword(t *testing.T) {
	client, err := NewClient("http://example.com")
	if err != nil {
		t.Errorf("Couldn't create client (%s)", err)
	}

	client.SetPassword("hello")

	has := false
	for _, cookie := range client.client.Jar.Cookies(client.baseUri) {
		has = cookie.Name == PASSWORD_COOKIE_NAME && cookie.Value == "hello"
		if has {
			break
		}
	}

	if !has {
		t.Errorf("Can not find necessary cookie")
	}
}
