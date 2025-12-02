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

	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// AddResponseErrorMiddleware adds response error wrapper middleware
func AddResponseErrorMiddleware(stack *middleware.Stack) error {
	// add error wrapper middleware before request id retriever middleware so that it can wrap the error response
	// returned by operation deserializers
	return stack.Deserialize.Insert(&ResponseErrorWrapper{}, "RequestIDRetriever", middleware.Before)
}

// ResponseErrorWrapper wraps operation errors with ResponseError.
type ResponseErrorWrapper struct {
}

// ID returns the middleware identifier
func (m *ResponseErrorWrapper) ID() string {
	return "ResponseErrorWrapper"
}

// HandleDeserialize wraps the stack error with smithyhttp.ResponseError.
func (m *ResponseErrorWrapper) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err == nil {
		// Nothing to do when there is no error.
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		// No raw response to wrap with.
		return out, metadata, err
	}

	// look for request id in metadata
	reqID, _ := awsmiddleware.GetRequestIDMetadata(metadata)

	// Wrap the returned smithy error with the request id retrieved from the metadata
	err = &ResponseError{
		ResponseError: &smithyhttp.ResponseError{
			Response: resp,
			Err:      err,
		},
		RequestID: reqID,
	}

	return out, metadata, err
}
