package serials_now_api

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	LOGIN_COOKIE_NAME    = "Login"
	PASSWORD_COOKIE_NAME = "Password"
)

type Client struct {
	baseUri *url.URL
	client  *http.Client
}

func NewClient(baseUri string) (*Client, error) {
	uri, err := url.Parse(baseUri)
	if err != nil {
		return &Client{}, err
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		baseUri: uri,
		client: &http.Client{
			Jar: jar,
		},
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

func (c *Client) SetLogin(login string) {
	c.SetCookie(LOGIN_COOKIE_NAME, login)
}

func (c *Client) SetPassword(password string) {
	c.SetCookie(PASSWORD_COOKIE_NAME, password)
}

func (c *Client) SetCookie(name, value string) {
	set := false

	cookies := c.client.Jar.Cookies(c.baseUri)
	for _, cookie := range cookies {
		if cookie.Name == name {
			cookie.Value = value
			set = true
			break
		}
	}

	if !set {
		cookie := &http.Cookie{
			Name:  name,
			Value: value,
		}

		cookies = append(cookies, cookie)
	}

	c.client.Jar.SetCookies(c.baseUri, cookies)
}
