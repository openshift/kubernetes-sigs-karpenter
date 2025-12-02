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

package http

import (
	"context"
	"io"

	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
)

// AddErrorCloseResponseBodyMiddleware adds the middleware to automatically
// close the response body of an operation request if the request response
// failed.
func AddErrorCloseResponseBodyMiddleware(stack *middleware.Stack) error {
	return stack.Deserialize.Insert(&errorCloseResponseBodyMiddleware{}, "OperationDeserializer", middleware.Before)
}

type errorCloseResponseBodyMiddleware struct{}

func (*errorCloseResponseBodyMiddleware) ID() string {
	return "ErrorCloseResponseBody"
}

func (m *errorCloseResponseBodyMiddleware) HandleDeserialize(
	ctx context.Context, input middleware.DeserializeInput, next middleware.DeserializeHandler,
) (
	output middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err := next.HandleDeserialize(ctx, input)
	if err != nil {
		if resp, ok := out.RawResponse.(*Response); ok && resp != nil && resp.Body != nil {
			// Consume the full body to prevent TCP connection resets on some platforms
			_, _ = io.Copy(io.Discard, resp.Body)
			// Do not validate that the response closes successfully.
			resp.Body.Close()
		}
	}

	return out, metadata, err
}

// AddCloseResponseBodyMiddleware adds the middleware to automatically close
// the response body of an operation request, after the response had been
// deserialized.
func AddCloseResponseBodyMiddleware(stack *middleware.Stack) error {
	return stack.Deserialize.Insert(&closeResponseBody{}, "OperationDeserializer", middleware.Before)
}

type closeResponseBody struct{}

func (*closeResponseBody) ID() string {
	return "CloseResponseBody"
}

func (m *closeResponseBody) HandleDeserialize(
	ctx context.Context, input middleware.DeserializeInput, next middleware.DeserializeHandler,
) (
	output middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err := next.HandleDeserialize(ctx, input)
	if err != nil {
		return out, metadata, err
	}

	if resp, ok := out.RawResponse.(*Response); ok {
		// Consume the full body to prevent TCP connection resets on some platforms
		_, copyErr := io.Copy(io.Discard, resp.Body)
		if copyErr != nil {
			middleware.GetLogger(ctx).Logf(logging.Warn, "failed to discard remaining HTTP response body, this may affect connection reuse")
		}

		closeErr := resp.Body.Close()
		if closeErr != nil {
			middleware.GetLogger(ctx).Logf(logging.Warn, "failed to close HTTP response body, this may affect connection reuse")
		}
	}

	return out, metadata, err
}
