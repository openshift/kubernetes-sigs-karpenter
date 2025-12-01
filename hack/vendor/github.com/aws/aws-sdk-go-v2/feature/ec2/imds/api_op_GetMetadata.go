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
	"fmt"
	"io"

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

const getMetadataPath = "/latest/meta-data"

// GetMetadata uses the path provided to request information from the Amazon
// EC2 Instance Metadata Service. The content will be returned as a string, or
// error if the request failed.
func (c *Client) GetMetadata(ctx context.Context, params *GetMetadataInput, optFns ...func(*Options)) (*GetMetadataOutput, error) {
	if params == nil {
		params = &GetMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetadata", params, optFns,
		addGetMetadataMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetMetadataInput provides the input parameters for the GetMetadata
// operation.
type GetMetadataInput struct {
	// The relative metadata path to retrieve. Can be empty string to retrieve
	// a response containing a new line separated list of metadata resources
	// available.
	//
	// Must not include the metadata base path.
	//
	// May include leading slash. If Path includes trailing slash the trailing slash
	// will be included in the request for the resource.
	Path string
}

// GetMetadataOutput provides the output parameters for the GetMetadata
// operation.
type GetMetadataOutput struct {
	Content io.ReadCloser

	ResultMetadata middleware.Metadata
}

func addGetMetadataMiddleware(stack *middleware.Stack, options Options) error {
	return addAPIRequestMiddleware(stack,
		options,
		"GetMetadata",
		buildGetMetadataPath,
		buildGetMetadataOutput)
}

func buildGetMetadataPath(params interface{}) (string, error) {
	p, ok := params.(*GetMetadataInput)
	if !ok {
		return "", fmt.Errorf("unknown parameter type %T", params)
	}

	return appendURIPath(getMetadataPath, p.Path), nil
}

func buildGetMetadataOutput(resp *smithyhttp.Response) (interface{}, error) {
	return &GetMetadataOutput{
		Content: resp.Body,
	}, nil
}
