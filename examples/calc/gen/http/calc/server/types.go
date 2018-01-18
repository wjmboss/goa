// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package server

import (
	goa "goa.design/goa"
	calcsvc "goa.design/goa/examples/calc/gen/calc"
)

// AddedRequestBody is the type of the "calc" service "added" endpoint HTTP
// request body.
type AddedRequestBody struct {
	Operands map[string]int `form:"operands,omitempty" json:"operands,omitempty" xml:"operands,omitempty"`
}

// NewAddAddPayload builds a calc service add endpoint payload.
func NewAddAddPayload(a int, b int) *calcsvc.AddPayload {
	return &calcsvc.AddPayload{
		A: a,
		B: b,
	}
}

// NewAddedAddedPayload builds a calc service added endpoint payload.
func NewAddedAddedPayload(body *AddedRequestBody) *calcsvc.AddedPayload {
	v := &calcsvc.AddedPayload{}
	v.Operands = make(map[string]int, len(body.Operands))
	for key, val := range body.Operands {
		tk := key
		tv := val
		v.Operands[tk] = tv
	}
	return v
}

// Validate runs the validations defined on AddedRequestBody
func (body *AddedRequestBody) Validate() (err error) {
	if body.Operands == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("operands", "body"))
	}
	return
}
