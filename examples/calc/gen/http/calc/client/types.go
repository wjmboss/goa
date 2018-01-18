// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP client types
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package client

import (
	calcsvc "goa.design/goa/examples/calc/gen/calc"
)

// AddedRequestBody is the type of the "calc" service "added" endpoint HTTP
// request body.
type AddedRequestBody struct {
	Operands map[string]int `form:"operands" json:"operands" xml:"operands"`
}

// NewAddedRequestBody builds the HTTP request body from the payload of the
// "added" endpoint of the "calc" service.
func NewAddedRequestBody(p *calcsvc.AddedPayload) *AddedRequestBody {
	body := &AddedRequestBody{}
	if p.Operands != nil {
		body.Operands = make(map[string]int, len(p.Operands))
		for key, val := range p.Operands {
			tk := key
			tv := val
			body.Operands[tk] = tv
		}
	}
	return body
}
