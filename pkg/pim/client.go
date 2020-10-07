// Package pimclient provides logic for setting up the API pimclient
package pim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// A Client manages communication with the PIM API.
type Client struct {

	// Base URL for API requests. BaseURL should
	// always be specified with a trailing slash.
	baseURL    *url.URL
	httpClient *http.Client // HTTP client used to communicate with the API.
	// User agent used when communicating with the PIM API.
	UserAgent string

	common service
	// Services used for talking to different parts of the PIM API.
	WarehouseLocations *WarehouseLocations
	Products           *Products
	Attributes         *Attributes
	Brands             *Brands
	Categories         *Categories
	Families           *Families
}

// NewClient returns a new PIM API client. If a nil httpClient is
// provided, a wrapper's default http.Client will be used.
//
// Deprecated: NewClient exists for historical compatibility
// and should not be used. To create the new client
// use the NewAPIClient with the user agent parameter.
func NewClient(baseURL *url.URL, httpCli *http.Client) *Client {
	c := &Client{
		baseURL:   baseURL,
		UserAgent: "pim-wrapper",
	}
	if httpCli != nil {
		c.httpClient = httpCli
	} else {
		c.httpClient = getDefaultHTTPClient()
	}
	c.common.client = c
	c.WarehouseLocations = (*WarehouseLocations)(&c.common)
	c.Products = (*Products)(&c.common)
	c.Attributes = (*Attributes)(&c.common)
	c.Brands = (*Brands)(&c.common)
	c.Categories = (*Categories)(&c.common)
	c.Families = (*Families)(&c.common)
	return c
}

// NewAPIClient returns a new PIM API client. If a nil httpClient is
// provided, a wrapper's default http.Client will be used.
func NewAPIClient(baseURL *url.URL, httpCli *http.Client, userAgent string) *Client {
	c := &Client{
		baseURL:   baseURL,
		UserAgent: userAgent,
	}
	if httpCli != nil {
		c.httpClient = httpCli
	} else {
		c.httpClient = getDefaultHTTPClient()
	}
	c.common.client = c
	c.WarehouseLocations = (*WarehouseLocations)(&c.common)
	c.Products = (*Products)(&c.common)
	c.Attributes = (*Attributes)(&c.common)
	c.Brands = (*Brands)(&c.common)
	c.Categories = (*Categories)(&c.common)
	c.Families = (*Families)(&c.common)
	return c
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context")
		default:
		}
		return nil, errors.Wrap(err, "client Do")
	}

	defer func() {
		// Ensure the response body is fully read and closed
		// before we reconnect, so that we reuse the same TCP connection.
		// Close the previous response's body. But read at least some of
		// the body so if it's small the underlying TCP connection will be
		// re-used. No need to check for errors: if it fails, the Transport
		// won't reuse it anyway.
		const maxBodySlurpSize = 2 << 10
		if resp.ContentLength == -1 || resp.ContentLength <= maxBodySlurpSize {
			if _, err := io.CopyN(ioutil.Discard, resp.Body, maxBodySlurpSize); err != nil {
				if err == io.EOF {
					err = nil // ignore EOF errors caused by empty response body
				} else {
					logrus.Error(err)
				}
			}
		}

		if err := resp.Body.Close(); err != nil {
			logrus.Errorf("%s when closing the response body", err)
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			if _, err := io.Copy(w, resp.Body); err != nil {
				return nil, errors.Wrap(err, "io.Copy failed reading the body")
			}
		} else {
			errResp := bytes.Buffer{}
			tee := io.TeeReader(resp.Body, &errResp)
			decErr := json.NewDecoder(tee).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				//if the response is not of expected structure perhaps that's an error response
				ev := &MessageResponse{}
				if err := json.NewDecoder(&errResp).Decode(ev); err != nil {
					return nil, err
				}
				return nil, errors.Wrap(errors.New(ev.Message), "got error response with message")
			}
		}
	}

	return resp, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.baseURL)
	}
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}
