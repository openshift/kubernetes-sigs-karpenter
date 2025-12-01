/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validate

import (
	"context"
)

// validateCtxKey is the key type of context key in this pkg
type validateCtxKey string

const (
	operationTypeKey validateCtxKey = "operationTypeKey"
)

type operationType string

const (
	request  operationType = "request"
	response operationType = "response"
	none     operationType = "none" // not specified in ctx
)

var operationTypeEnum = []operationType{request, response, none}

// WithOperationRequest returns a new context with operationType request
// in context value
func WithOperationRequest(ctx context.Context) context.Context {
	return withOperation(ctx, request)
}

// WithOperationRequest returns a new context with operationType response
// in context value
func WithOperationResponse(ctx context.Context) context.Context {
	return withOperation(ctx, response)
}

func withOperation(ctx context.Context, operation operationType) context.Context {
	return context.WithValue(ctx, operationTypeKey, operation)
}

// extractOperationType extracts the operation type from ctx
// if not specified or of unknown value, return none operation type
func extractOperationType(ctx context.Context) operationType {
	v := ctx.Value(operationTypeKey)
	if v == nil {
		return none
	}
	res, ok := v.(operationType)
	if !ok {
		return none
	}
	// validate the value is in operation enum
	if err := Enum("", "", res, operationTypeEnum); err != nil {
		return none
	}
	return res
}
