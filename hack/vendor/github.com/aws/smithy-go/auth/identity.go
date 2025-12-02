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

package auth

import (
	"context"
	"time"

	"github.com/aws/smithy-go"
)

// Identity contains information that identifies who the user making the
// request is.
type Identity interface {
	Expiration() time.Time
}

// IdentityResolver defines the interface through which an Identity is
// retrieved.
type IdentityResolver interface {
	GetIdentity(context.Context, smithy.Properties) (Identity, error)
}

// IdentityResolverOptions defines the interface through which an entity can be
// queried to retrieve an IdentityResolver for a given auth scheme.
type IdentityResolverOptions interface {
	GetIdentityResolver(schemeID string) IdentityResolver
}

// AnonymousIdentity is a sentinel to indicate no identity.
type AnonymousIdentity struct{}

var _ Identity = (*AnonymousIdentity)(nil)

// Expiration returns the zero value for time, as anonymous identity never
// expires.
func (*AnonymousIdentity) Expiration() time.Time {
	return time.Time{}
}

// AnonymousIdentityResolver returns AnonymousIdentity.
type AnonymousIdentityResolver struct{}

var _ IdentityResolver = (*AnonymousIdentityResolver)(nil)

// GetIdentity returns AnonymousIdentity.
func (*AnonymousIdentityResolver) GetIdentity(_ context.Context, _ smithy.Properties) (Identity, error) {
	return &AnonymousIdentity{}, nil
}
