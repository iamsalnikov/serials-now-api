package serial

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var UnexpectedStatusCode = errors.New("unexpected status code")

type Endpoint struct {
	SerialID int64
	Data     EndpointData
}

func NewEndpoint(serialID int64) *Endpoint {
	return &Endpoint{
		SerialID: serialID,
	}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	form := url.Values{}
	form.Set("ID", strconv.FormatInt(e.SerialID, 10))

	request, err := http.NewRequest(http.MethodPost, "/API/Online.php", strings.NewReader(form.Encode()))
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

	decoder := json.NewDecoder(response.Body)

	return decoder.Decode(&e.Data)
}
