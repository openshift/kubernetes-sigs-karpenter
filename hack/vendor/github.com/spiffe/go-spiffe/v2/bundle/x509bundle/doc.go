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

// Package x509bundle provides X.509 bundle related functionality.
//
// A bundle represents a collection of X.509 authorities, i.e., those that
// are used to authenticate SPIFFE X509-SVIDs.
//
// You can create a new bundle for a specific trust domain:
//
//	td := spiffeid.RequireTrustDomainFromString("example.org")
//	bundle := x509bundle.New(td)
//
// Or you can load it from disk:
//
//	td := spiffeid.RequireTrustDomainFromString("example.org")
//	bundle := x509bundle.Load(td, "bundle.pem")
//
// The bundle can be initialized with X.509 authorities:
//
//	td := spiffeid.RequireTrustDomainFromString("example.org")
//	var x509Authorities []*x509.Certificate = ...
//	bundle := x509bundle.FromX509Authorities(td, x509Authorities)
//
// In addition, you can add X.509 authorities to the bundle:
//
//	var x509CA *x509.Certificate = ...
//	bundle.AddX509Authority(x509CA)
//
// Bundles can be organized into a set, keyed by trust domain:
//
//	set := x509bundle.NewSet()
//	set.Add(bundle)
//
// A Source is source of X.509 bundles for a trust domain. Both the Bundle
// and Set types implement Source:
//
//	// Initialize the source from a bundle or set
//	var source x509bundle.Source = bundle
//	// ... or ...
//	var source x509bundle.Source = set
//
//	// Use the source to query for bundles by trust domain
//	bundle, err := source.GetX509BundleForTrustDomain(td)
package x509bundle
