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

// AuthScheme defines an HTTP authentication scheme.
type AuthScheme interface {
	SchemeID() string
	IdentityResolver(auth.IdentityResolverOptions) auth.IdentityResolver
	Signer() Signer
}

// Signer defines the interface through which HTTP requests are supplemented
// with an Identity.
type Signer interface {
	SignRequest(context.Context, *Request, auth.Identity, smithy.Properties) error
}
