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

package smithy

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/auth/bearer"
)

// BearerTokenAdapter adapts smithy bearer.Token to smithy auth.Identity.
type BearerTokenAdapter struct {
	Token bearer.Token
}

var _ auth.Identity = (*BearerTokenAdapter)(nil)

// Expiration returns the time of expiration for the token.
func (v *BearerTokenAdapter) Expiration() time.Time {
	return v.Token.Expires
}

// BearerTokenProviderAdapter adapts smithy bearer.TokenProvider to smithy
// auth.IdentityResolver.
type BearerTokenProviderAdapter struct {
	Provider bearer.TokenProvider
}

var _ (auth.IdentityResolver) = (*BearerTokenProviderAdapter)(nil)

// GetIdentity retrieves a bearer token using the underlying provider.
func (v *BearerTokenProviderAdapter) GetIdentity(ctx context.Context, _ smithy.Properties) (
	auth.Identity, error,
) {
	token, err := v.Provider.RetrieveBearerToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("get token: %w", err)
	}

	return &BearerTokenAdapter{Token: token}, nil
}
