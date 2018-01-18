// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa"
	calcsvc "goa.design/goa/examples/calc/gen/calc"
)

// BuildAddAddPayload builds the payload for the calc add endpoint from CLI
// flags.
func BuildAddAddPayload(calcAddA string, calcAddB string) (*calcsvc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddA, 10, 64)
		a = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddB, 10, 64)
		b = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for b, must be INT")
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &calcsvc.AddPayload{
		A: a,
		B: b,
	}
	return payload, nil
}

// BuildAddedAddedPayload builds the payload for the calc added endpoint from
// CLI flags.
func BuildAddedAddedPayload(calcAddedBody string) (*calcsvc.AddedPayload, error) {
	var err error
	var body AddedRequestBody
	{
		err = json.Unmarshal([]byte(calcAddedBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"operands\": {\n         \"Ipsam voluptatem.\": 3237209857320107068\n      }\n   }'")
		}
		if body.Operands == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("operands", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	v := &calcsvc.AddedPayload{}
	if body.Operands != nil {
		v.Operands = make(map[string]int, len(body.Operands))
		for key, val := range body.Operands {
			tk := key
			tv := val
			v.Operands[tk] = tv
		}
	}
	return v, nil
}
