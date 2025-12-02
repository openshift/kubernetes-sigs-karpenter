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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
)

// CredentialsAdapter adapts aws.Credentials to auth.Identity.
type CredentialsAdapter struct {
	Credentials aws.Credentials
}

var _ auth.Identity = (*CredentialsAdapter)(nil)

// Expiration returns the time of expiration for the credentials.
func (v *CredentialsAdapter) Expiration() time.Time {
	return v.Credentials.Expires
}

// CredentialsProviderAdapter adapts aws.CredentialsProvider to auth.IdentityResolver.
type CredentialsProviderAdapter struct {
	Provider aws.CredentialsProvider
}

var _ (auth.IdentityResolver) = (*CredentialsProviderAdapter)(nil)

// GetIdentity retrieves AWS credentials using the underlying provider.
func (v *CredentialsProviderAdapter) GetIdentity(ctx context.Context, _ smithy.Properties) (
	auth.Identity, error,
) {
	if v.Provider == nil {
		return &CredentialsAdapter{Credentials: aws.Credentials{}}, nil
	}

	creds, err := v.Provider.Retrieve(ctx)
	if err != nil {
		return nil, fmt.Errorf("get credentials: %w", err)
	}

	return &CredentialsAdapter{Credentials: creds}, nil
}
