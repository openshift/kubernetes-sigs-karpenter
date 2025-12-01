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

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// GetRegion retrieves an identity document describing an
// instance. Error is returned if the request fails or is unable to parse
// the response.
func (c *Client) GetRegion(
	ctx context.Context, params *GetRegionInput, optFns ...func(*Options),
) (
	*GetRegionOutput, error,
) {
	if params == nil {
		params = &GetRegionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRegion", params, optFns,
		addGetRegionMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRegionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetRegionInput provides the input parameters for GetRegion operation.
type GetRegionInput struct{}

// GetRegionOutput provides the output parameters for GetRegion operation.
type GetRegionOutput struct {
	Region string

	ResultMetadata middleware.Metadata
}

func addGetRegionMiddleware(stack *middleware.Stack, options Options) error {
	return addAPIRequestMiddleware(stack,
		options,
		"GetRegion",
		buildGetInstanceIdentityDocumentPath,
		buildGetRegionOutput,
	)
}

func buildGetRegionOutput(resp *smithyhttp.Response) (interface{}, error) {
	out, err := buildGetInstanceIdentityDocumentOutput(resp)
	if err != nil {
		return nil, err
	}

	result, ok := out.(*GetInstanceIdentityDocumentOutput)
	if !ok {
		return nil, fmt.Errorf("unexpected instance identity document type, %T", out)
	}

	region := result.Region
	if len(region) == 0 {
		return "", fmt.Errorf("instance metadata did not return a region value")
	}

	return &GetRegionOutput{
		Region: region,
	}, nil
}
