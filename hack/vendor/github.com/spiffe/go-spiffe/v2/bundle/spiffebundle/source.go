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

package spiffebundle

import "github.com/spiffe/go-spiffe/v2/spiffeid"

// Source represents a source of SPIFFE bundles keyed by trust domain.
type Source interface {
	// GetBundleForTrustDomain returns the SPIFFE bundle for the given trust
	// domain.
	GetBundleForTrustDomain(trustDomain spiffeid.TrustDomain) (*Bundle, error)
}
