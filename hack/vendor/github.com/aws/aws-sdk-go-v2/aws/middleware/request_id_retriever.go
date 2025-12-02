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

	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// AddRequestIDRetrieverMiddleware adds request id retriever middleware
func AddRequestIDRetrieverMiddleware(stack *middleware.Stack) error {
	// add error wrapper middleware before operation deserializers so that it can wrap the error response
	// returned by operation deserializers
	return stack.Deserialize.Insert(&RequestIDRetriever{}, "OperationDeserializer", middleware.Before)
}

// RequestIDRetriever middleware captures the AWS service request ID from the
// raw response.
type RequestIDRetriever struct {
}

// ID returns the middleware identifier
func (m *RequestIDRetriever) ID() string {
	return "RequestIDRetriever"
}

// HandleDeserialize pulls the AWS request ID from the response, storing it in
// operation metadata.
func (m *RequestIDRetriever) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		// No raw response to wrap with.
		return out, metadata, err
	}

	// Different header which can map to request id
	requestIDHeaderList := []string{"X-Amzn-Requestid", "X-Amz-RequestId"}

	for _, h := range requestIDHeaderList {
		// check for headers known to contain Request id
		if v := resp.Header.Get(h); len(v) != 0 {
			// set reqID on metadata for successful responses.
			SetRequestIDMetadata(&metadata, v)

			span, _ := tracing.GetSpan(ctx)
			span.SetProperty("aws.request_id", v)
			break
		}
	}

	return out, metadata, err
}
