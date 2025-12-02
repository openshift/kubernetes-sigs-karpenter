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

package swarm

import (
	"strconv"
	"time"
)

// Version represents the internal object version.
type Version struct {
	Index uint64 `json:",omitempty"`
}

// String implements fmt.Stringer interface.
func (v Version) String() string {
	return strconv.FormatUint(v.Index, 10)
}

// Meta is a base object inherited by most of the other once.
type Meta struct {
	Version   Version   `json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

// Annotations represents how to describe an object.
type Annotations struct {
	Name   string            `json:",omitempty"`
	Labels map[string]string `json:"Labels"`
}

// Driver represents a driver (network, logging, secrets backend).
type Driver struct {
	Name    string            `json:",omitempty"`
	Options map[string]string `json:",omitempty"`
}

// TLSInfo represents the TLS information about what CA certificate is trusted,
// and who the issuer for a TLS certificate is
type TLSInfo struct {
	// TrustRoot is the trusted CA root certificate in PEM format
	TrustRoot string `json:",omitempty"`

	// CertIssuer is the raw subject bytes of the issuer
	CertIssuerSubject []byte `json:",omitempty"`

	// CertIssuerPublicKey is the raw public key bytes of the issuer
	CertIssuerPublicKey []byte `json:",omitempty"`
}
