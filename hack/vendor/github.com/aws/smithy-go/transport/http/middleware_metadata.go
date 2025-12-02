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

	"github.com/aws/smithy-go/middleware"
)

type (
	hostnameImmutableKey struct{}
	hostPrefixDisableKey struct{}
)

// GetHostnameImmutable retrieves whether the endpoint hostname should be considered
// immutable or not.
//
// Scoped to stack values. Use middleware#ClearStackValues to clear all stack
// values.
func GetHostnameImmutable(ctx context.Context) (v bool) {
	v, _ = middleware.GetStackValue(ctx, hostnameImmutableKey{}).(bool)
	return v
}

// SetHostnameImmutable sets or modifies whether the request's endpoint hostname
// should be considered immutable or not.
//
// Scoped to stack values. Use middleware#ClearStackValues to clear all stack
// values.
func SetHostnameImmutable(ctx context.Context, value bool) context.Context {
	return middleware.WithStackValue(ctx, hostnameImmutableKey{}, value)
}

// IsEndpointHostPrefixDisabled retrieves whether the hostname prefixing is
// disabled.
//
// Scoped to stack values. Use middleware#ClearStackValues to clear all stack
// values.
func IsEndpointHostPrefixDisabled(ctx context.Context) (v bool) {
	v, _ = middleware.GetStackValue(ctx, hostPrefixDisableKey{}).(bool)
	return v
}

// DisableEndpointHostPrefix sets or modifies whether the request's endpoint host
// prefixing should be disabled. If value is true, endpoint host prefixing
// will be disabled.
//
// Scoped to stack values. Use middleware#ClearStackValues to clear all stack
// values.
func DisableEndpointHostPrefix(ctx context.Context, value bool) context.Context {
	return middleware.WithStackValue(ctx, hostPrefixDisableKey{}, value)
}
