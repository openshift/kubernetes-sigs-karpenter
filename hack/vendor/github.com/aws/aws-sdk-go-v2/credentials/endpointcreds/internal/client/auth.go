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

package client

import (
	"context"
	"github.com/aws/smithy-go/middleware"
)

type getIdentityMiddleware struct {
	options Options
}

func (*getIdentityMiddleware) ID() string {
	return "GetIdentity"
}

func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	return next.HandleFinalize(ctx, in)
}

type signRequestMiddleware struct {
}

func (*signRequestMiddleware) ID() string {
	return "Signing"
}

func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	return next.HandleFinalize(ctx, in)
}

type resolveAuthSchemeMiddleware struct {
	operation string
	options   Options
}

func (*resolveAuthSchemeMiddleware) ID() string {
	return "ResolveAuthScheme"
}

func (m *resolveAuthSchemeMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	return next.HandleFinalize(ctx, in)
}
