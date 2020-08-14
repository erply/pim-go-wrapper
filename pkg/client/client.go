// Package client provides logic for setting up the API client
package client

import (
	"net/http"
)

type Client struct {
	Url, sessionKey, clientCode string
	httpClient                  *http.Client
}

func NewClient(sk, cc, url string, httpCli *http.Client) *Client {
	c := &Client{
		Url:        url,
		sessionKey: sk,
		clientCode: cc,
	}
	if httpCli != nil {
		c.httpClient = httpCli
	} else {
		c.httpClient = getDefaultHTTPClient()
	}
	return c
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("clientCode", c.clientCode)
	req.Header.Set("sessionKey", c.sessionKey)
	resp, err := c.httpClient.Do(req)
	return resp, err
}
