package serials_now_api

import (
	"net/http"
	"fmt"
	"os"
	"net/url"
)

type Client struct {
	baseUri *url.URL
	client *http.Client
}

func NewClient(baseUri string) (*Client, error) {
	uri, err := url.Parse(baseUri)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		baseUri: uri,
		client: http.DefaultClient,
	}, nil
}

func (c *Client) Send(endpoint EndpointInterface) error {
	req, err := endpoint.BuildHttpRequest()
	if err != nil {
		return err
	}

	req.Host = c.baseUri.Host
	req.URL.Scheme = c.baseUri.Scheme
	req.URL.Host = c.baseUri.Host

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return endpoint.ParseResponse(resp)
}