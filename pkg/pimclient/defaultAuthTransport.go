package pimclient

import "net/http"

// DefaultAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Authentication with the provided session key and pimclient code.
type DefaultAuthTransport struct {
	SessionKey string // ERPLY sessionKey
	ClientCode string // ERPLY clientCode

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func NewDefaultAuthTransport(sessionKey, clientCode string, transport http.RoundTripper) *DefaultAuthTransport {
	return &DefaultAuthTransport{SessionKey: sessionKey, ClientCode: clientCode, Transport: transport}
}

// RoundTrip implements the RoundTripper interface.
func (t *DefaultAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("clientCode", t.ClientCode)
	req.Header.Add("sessionKey", t.SessionKey)
	return t.transport().RoundTrip(req)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *DefaultAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}
func (t *DefaultAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
