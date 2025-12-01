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
	"strconv"
	"strings"
	"time"

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

const getTokenPath = "/latest/api/token"
const tokenTTLHeader = "X-Aws-Ec2-Metadata-Token-Ttl-Seconds"

// getToken uses the duration to return a token for EC2 IMDS, or an error if
// the request failed.
func (c *Client) getToken(ctx context.Context, params *getTokenInput, optFns ...func(*Options)) (*getTokenOutput, error) {
	if params == nil {
		params = &getTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "getToken", params, optFns,
		addGetTokenMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*getTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type getTokenInput struct {
	TokenTTL time.Duration
}

type getTokenOutput struct {
	Token    string
	TokenTTL time.Duration

	ResultMetadata middleware.Metadata
}

func addGetTokenMiddleware(stack *middleware.Stack, options Options) error {
	err := addRequestMiddleware(stack,
		options,
		"PUT",
		"GetToken",
		buildGetTokenPath,
		buildGetTokenOutput)
	if err != nil {
		return err
	}

	err = stack.Serialize.Add(&tokenTTLRequestHeader{}, middleware.After)
	if err != nil {
		return err
	}

	return nil
}

func buildGetTokenPath(interface{}) (string, error) {
	return getTokenPath, nil
}

func buildGetTokenOutput(resp *smithyhttp.Response) (v interface{}, err error) {
	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		} else if closeErr != nil {
			err = fmt.Errorf("response body close error: %v, original error: %w", closeErr, err)
		}
	}()

	ttlHeader := resp.Header.Get(tokenTTLHeader)
	tokenTTL, err := strconv.ParseInt(ttlHeader, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("unable to parse API token, %w", err)
	}

	var token strings.Builder
	if _, err = io.Copy(&token, resp.Body); err != nil {
		return nil, fmt.Errorf("unable to read API token, %w", err)
	}

	return &getTokenOutput{
		Token:    token.String(),
		TokenTTL: time.Duration(tokenTTL) * time.Second,
	}, nil
}

type tokenTTLRequestHeader struct{}

func (*tokenTTLRequestHeader) ID() string { return "tokenTTLRequestHeader" }
func (*tokenTTLRequestHeader) HandleSerialize(
	ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("expect HTTP transport, got %T", in.Request)
	}

	input, ok := in.Parameters.(*getTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("expect getTokenInput, got %T", in.Parameters)
	}

	req.Header.Set(tokenTTLHeader, strconv.Itoa(int(input.TokenTTL/time.Second)))

	return next.HandleSerialize(ctx, in)
}
