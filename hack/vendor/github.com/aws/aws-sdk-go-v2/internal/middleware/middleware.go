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

package middleware

import (
	"context"
	"sync/atomic"
	"time"

	internalcontext "github.com/aws/aws-sdk-go-v2/internal/context"
	"github.com/aws/smithy-go/middleware"
)

// AddTimeOffsetMiddleware sets a value representing clock skew on the request context.
// This can be read by other operations (such as signing) to correct the date value they send
// on the request
type AddTimeOffsetMiddleware struct {
	Offset *atomic.Int64
}

// ID the identifier for AddTimeOffsetMiddleware
func (m *AddTimeOffsetMiddleware) ID() string { return "AddTimeOffsetMiddleware" }

// HandleBuild sets a value for attemptSkew on the request context if one is set on the client.
func (m AddTimeOffsetMiddleware) HandleBuild(ctx context.Context, in middleware.BuildInput, next middleware.BuildHandler) (
	out middleware.BuildOutput, metadata middleware.Metadata, err error,
) {
	if m.Offset != nil {
		offset := time.Duration(m.Offset.Load())
		ctx = internalcontext.SetAttemptSkewContext(ctx, offset)
	}
	return next.HandleBuild(ctx, in)
}

// HandleDeserialize gets the clock skew context from the context, and if set, sets it on the pointer
// held by AddTimeOffsetMiddleware
func (m *AddTimeOffsetMiddleware) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	if v := internalcontext.GetAttemptSkewContext(ctx); v != 0 {
		m.Offset.Store(v.Nanoseconds())
	}
	return next.HandleDeserialize(ctx, in)
}
