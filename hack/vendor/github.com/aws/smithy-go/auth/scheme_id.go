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

// Anonymous
const (
	SchemeIDAnonymous = "smithy.api#noAuth"
)

// HTTP auth schemes
const (
	SchemeIDHTTPBasic  = "smithy.api#httpBasicAuth"
	SchemeIDHTTPDigest = "smithy.api#httpDigestAuth"
	SchemeIDHTTPBearer = "smithy.api#httpBearerAuth"
	SchemeIDHTTPAPIKey = "smithy.api#httpApiKeyAuth"
)

// AWS auth schemes
const (
	SchemeIDSigV4  = "aws.auth#sigv4"
	SchemeIDSigV4A = "aws.auth#sigv4a"
)
