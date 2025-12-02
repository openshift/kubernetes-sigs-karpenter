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

	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
)

// NewAnonymousScheme returns the anonymous HTTP auth scheme.
func NewAnonymousScheme() AuthScheme {
	return &authScheme{
		schemeID: auth.SchemeIDAnonymous,
		signer:   &nopSigner{},
	}
}

// authScheme is parameterized to generically implement the exported AuthScheme
// interface
type authScheme struct {
	schemeID string
	signer   Signer
}

var _ AuthScheme = (*authScheme)(nil)

func (s *authScheme) SchemeID() string {
	return s.schemeID
}

func (s *authScheme) IdentityResolver(o auth.IdentityResolverOptions) auth.IdentityResolver {
	return o.GetIdentityResolver(s.schemeID)
}

func (s *authScheme) Signer() Signer {
	return s.signer
}

type nopSigner struct{}

var _ Signer = (*nopSigner)(nil)

func (*nopSigner) SignRequest(context.Context, *Request, auth.Identity, smithy.Properties) error {
	return nil
}
