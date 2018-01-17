// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc client HTTP transport
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Added Doer is the HTTP client used to make requests to the added endpoint.
	AddedDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddedDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Added returns a endpoint that makes HTTP requests to the calc service added
// server.
func (c *Client) Added() goa.Endpoint {
	var (
		decodeResponse = DecodeAddedResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddedRequest(v)
		if err != nil {
			return nil, err
		}

		resp, err := c.AddedDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "added", err)
		}
		return decodeResponse(resp)
	}
}
