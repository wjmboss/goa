// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/http"
)

// EncodeAddedResponse returns an encoder for responses returned by the calc
// added endpoint.
func EncodeAddedResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddedRequest returns a decoder for requests sent to the calc added
// endpoint.
func DecodeAddedRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {

		return body, nil
	}
}
