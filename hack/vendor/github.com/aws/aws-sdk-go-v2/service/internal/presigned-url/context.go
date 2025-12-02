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

package presignedurl

import (
	"context"

	"github.com/aws/smithy-go/middleware"
)

// WithIsPresigning adds the isPresigning sentinel value to a context to signal
// that the middleware stack is using the presign flow.
//
// Scoped to stack values. Use github.com/aws/smithy-go/middleware#ClearStackValues
// to clear all stack values.
func WithIsPresigning(ctx context.Context) context.Context {
	return middleware.WithStackValue(ctx, isPresigningKey{}, true)
}

// GetIsPresigning returns if the context contains the isPresigning sentinel
// value for presigning flows.
//
// Scoped to stack values. Use github.com/aws/smithy-go/middleware#ClearStackValues
// to clear all stack values.
func GetIsPresigning(ctx context.Context) bool {
	v, _ := middleware.GetStackValue(ctx, isPresigningKey{}).(bool)
	return v
}

type isPresigningKey struct{}

// AddAsIsPresigningMiddleware adds a middleware to the head of the stack that
// will update the stack's context to be flagged as being invoked for the
// purpose of presigning.
func AddAsIsPresigningMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(asIsPresigningMiddleware{}, middleware.Before)
}

// AddAsIsPresigingMiddleware is an alias for backwards compatibility.
//
// Deprecated: This API was released with a typo. Use
// [AddAsIsPresigningMiddleware] instead.
func AddAsIsPresigingMiddleware(stack *middleware.Stack) error {
	return AddAsIsPresigningMiddleware(stack)
}

type asIsPresigningMiddleware struct{}

func (asIsPresigningMiddleware) ID() string { return "AsIsPresigningMiddleware" }

func (asIsPresigningMiddleware) HandleInitialize(
	ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler,
) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	ctx = WithIsPresigning(ctx)
	return next.HandleInitialize(ctx, in)
}
