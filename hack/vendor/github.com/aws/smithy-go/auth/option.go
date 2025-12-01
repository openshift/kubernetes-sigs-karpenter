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

import "github.com/aws/smithy-go"

type (
	authOptionsKey struct{}
)

// Option represents a possible authentication method for an operation.
type Option struct {
	SchemeID           string
	IdentityProperties smithy.Properties
	SignerProperties   smithy.Properties
}

// GetAuthOptions gets auth Options from Properties.
func GetAuthOptions(p *smithy.Properties) ([]*Option, bool) {
	v, ok := p.Get(authOptionsKey{}).([]*Option)
	return v, ok
}

// SetAuthOptions sets auth Options on Properties.
func SetAuthOptions(p *smithy.Properties, options []*Option) {
	p.Set(authOptionsKey{}, options)
}
