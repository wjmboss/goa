// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package calcsvc

import (
	"context"
)

// The calc service performs operations on numbers
type Service interface {
	// Added implements added.
	Added(context.Context, map[string]string) (int, error)
}
