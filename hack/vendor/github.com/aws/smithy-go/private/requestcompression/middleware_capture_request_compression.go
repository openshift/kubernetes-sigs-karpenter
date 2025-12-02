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

package requestcompression

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"net/http"
)

const captureUncompressedRequestID = "CaptureUncompressedRequest"

// AddCaptureUncompressedRequestMiddleware captures http request before compress encoding for check
func AddCaptureUncompressedRequestMiddleware(stack *middleware.Stack, buf *bytes.Buffer) error {
	return stack.Serialize.Insert(&captureUncompressedRequestMiddleware{
		buf: buf,
	}, "RequestCompression", middleware.Before)
}

type captureUncompressedRequestMiddleware struct {
	req   *http.Request
	buf   *bytes.Buffer
	bytes []byte
}

// ID returns id of the captureUncompressedRequestMiddleware
func (*captureUncompressedRequestMiddleware) ID() string {
	return captureUncompressedRequestID
}

// HandleSerialize captures request payload before it is compressed by request compression middleware
func (m *captureUncompressedRequestMiddleware) HandleSerialize(ctx context.Context, input middleware.SerializeInput, next middleware.SerializeHandler,
) (
	output middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := input.Request.(*smithyhttp.Request)
	if !ok {
		return output, metadata, fmt.Errorf("error when retrieving http request")
	}

	_, err = io.Copy(m.buf, request.GetStream())
	if err != nil {
		return output, metadata, fmt.Errorf("error when copying http request stream: %q", err)
	}
	if err = request.RewindStream(); err != nil {
		return output, metadata, fmt.Errorf("error when rewinding request stream: %q", err)
	}

	return next.HandleSerialize(ctx, input)
}
