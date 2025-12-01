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

package imds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

const getIAMInfoPath = getMetadataPath + "/iam/info"

// GetIAMInfo retrieves an identity document describing an
// instance. Error is returned if the request fails or is unable to parse
// the response.
func (c *Client) GetIAMInfo(
	ctx context.Context, params *GetIAMInfoInput, optFns ...func(*Options),
) (
	*GetIAMInfoOutput, error,
) {
	if params == nil {
		params = &GetIAMInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIAMInfo", params, optFns,
		addGetIAMInfoMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIAMInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetIAMInfoInput provides the input parameters for GetIAMInfo operation.
type GetIAMInfoInput struct{}

// GetIAMInfoOutput provides the output parameters for GetIAMInfo operation.
type GetIAMInfoOutput struct {
	IAMInfo

	ResultMetadata middleware.Metadata
}

func addGetIAMInfoMiddleware(stack *middleware.Stack, options Options) error {
	return addAPIRequestMiddleware(stack,
		options,
		"GetIAMInfo",
		buildGetIAMInfoPath,
		buildGetIAMInfoOutput,
	)
}

func buildGetIAMInfoPath(params interface{}) (string, error) {
	return getIAMInfoPath, nil
}

func buildGetIAMInfoOutput(resp *smithyhttp.Response) (v interface{}, err error) {
	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		} else if closeErr != nil {
			err = fmt.Errorf("response body close error: %v, original error: %w", closeErr, err)
		}
	}()

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	body := io.TeeReader(resp.Body, ringBuffer)

	imdsResult := &GetIAMInfoOutput{}
	if err = json.NewDecoder(body).Decode(&imdsResult.IAMInfo); err != nil {
		return nil, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode instance identity document, %w", err),
			Snapshot: ringBuffer.Bytes(),
		}
	}
	// Any code other success is an error
	if !strings.EqualFold(imdsResult.Code, "success") {
		return nil, fmt.Errorf("failed to get EC2 IMDS IAM info, %s",
			imdsResult.Code)
	}

	return imdsResult, nil
}

// IAMInfo provides the shape for unmarshaling an IAM info from the metadata
// API.
type IAMInfo struct {
	Code               string
	LastUpdated        time.Time
	InstanceProfileArn string
	InstanceProfileID  string
}
