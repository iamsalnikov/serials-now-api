package subscriptions

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var UnexpectedStatusCode = errors.New("unexpected status code")
var BadRequest = errors.New("bad request")

type Endpoint struct {
	SerialID     int64
	TranslatorID int64
}

func NewEndpoint(serialID, translatorID int64) *Endpoint {
	return &Endpoint{
		SerialID:     serialID,
		TranslatorID: translatorID,
	}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	form := &url.Values{}
	form.Set("ID", strconv.FormatInt(e.SerialID, 10))
	form.Set("T", strconv.FormatInt(e.TranslatorID, 10))

	request, err := http.NewRequest(http.MethodPost, "/API/Subscription.php", strings.NewReader(form.Encode()))
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
	if err != nil {
		return err
	}

	if answer.Response != "" {
		return BadRequest
	}

	return nil
}
