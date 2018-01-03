package registration

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var UnexpectedStatusCode = errors.New("unexpected status code")
var BadRequest = errors.New("bad request")

type Endpoint struct {
	Email    string
	Login    string
	Password string
}

func NewEndpoint(email, login, password string) *Endpoint {
	return &Endpoint{email, login, password}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	form := &url.Values{}
	form.Set("Email", e.Email)
	form.Set("Login", e.Login)
	form.Set("Password", e.Password)

	request, err := http.NewRequest(http.MethodPost, "/API/Registration.php", strings.NewReader(form.Encode()))
	if err != nil {
		return request, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	return request, nil
}

func (e *Endpoint) ParseResponse(response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return UnexpectedStatusCode
	}

	answer := &answer{}
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(answer)
	if err != nil && err != io.EOF {
		return err
	}

	if answer.Response != "" {
		return BadRequest
	}

	return nil
}
