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

	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/auth/bearer"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// BearerTokenSignerAdapter adapts smithy bearer.Signer to smithy http
// auth.Signer.
type BearerTokenSignerAdapter struct {
	Signer bearer.Signer
}

var _ (smithyhttp.Signer) = (*BearerTokenSignerAdapter)(nil)

// SignRequest signs the request with the provided bearer token.
func (v *BearerTokenSignerAdapter) SignRequest(ctx context.Context, r *smithyhttp.Request, identity auth.Identity, _ smithy.Properties) error {
	ca, ok := identity.(*BearerTokenAdapter)
	if !ok {
		return fmt.Errorf("unexpected identity type: %T", identity)
	}

	signed, err := v.Signer.SignWithBearerToken(ctx, ca.Token, r)
	if err != nil {
		return fmt.Errorf("sign request: %w", err)
	}

	*r = *signed.(*smithyhttp.Request)
	return nil
}
