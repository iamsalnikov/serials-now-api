package serials_now_api

import "net/http"

type EndpointInterface interface {
	BuildHttpRequest() (*http.Request, error)
	ParseResponse(*http.Response) error
}
