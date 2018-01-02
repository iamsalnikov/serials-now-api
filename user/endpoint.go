package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var UnexpectedStatusCode = errors.New("unexpected status code")
var UnexpectedAnswerLength = errors.New("unexpected answer length")

type Endpoint struct {
	Data UserData
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) BuildHttpRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/API/User.php", nil)
}

func (e *Endpoint) ParseResponse(response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return UnexpectedStatusCode
	}

	answer := make([]json.RawMessage, 0)
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&answer)
	if err != nil {
		return err
	}

	if len(answer) != 6 {
		return UnexpectedAnswerLength
	}

	return e.decodeUserData(answer)
}

func (e *Endpoint) decodeUserData(data []json.RawMessage) error {
	decoder := json.NewDecoder(bytes.NewBuffer(data[0]))
	err := decoder.Decode(&e.Data.FavoriteEpisodes)
	if err != nil {
		return err
	}

	decoder = json.NewDecoder(bytes.NewBuffer(data[1]))
	err = decoder.Decode(&e.Data.WatchedEpisodes)
	if err != nil {
		return err
	}

	decoder = json.NewDecoder(bytes.NewBuffer(data[2]))
	err = decoder.Decode(&e.Data.NewEpisodes)
	if err != nil {
		return err
	}

	decoder = json.NewDecoder(bytes.NewBuffer(data[3]))
	err = decoder.Decode(&e.Data.SubscribedSerials)
	if err != nil {
		return err
	}

	decoder = json.NewDecoder(bytes.NewBuffer(data[4]))
	err = decoder.Decode(&e.Data.VotedSerials)
	if err != nil {
		return err
	}

	decoder = json.NewDecoder(bytes.NewBuffer(data[5]))
	err = decoder.Decode(&e.Data.Comments)
	if err != nil {
		return err
	}

	return nil
}
