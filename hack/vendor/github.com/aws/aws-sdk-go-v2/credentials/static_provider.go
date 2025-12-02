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

package credentials

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	// StaticCredentialsName provides a name of Static provider
	StaticCredentialsName = "StaticCredentials"
)

// StaticCredentialsEmptyError is emitted when static credentials are empty.
type StaticCredentialsEmptyError struct{}

func (*StaticCredentialsEmptyError) Error() string {
	return "static credentials are empty"
}

// A StaticCredentialsProvider is a set of credentials which are set, and will
// never expire.
type StaticCredentialsProvider struct {
	Value aws.Credentials
	// These values are for reporting purposes and are not meant to be set up directly
	Source []aws.CredentialSource
}

// ProviderSources returns the credential chain that was used to construct this provider
func (s StaticCredentialsProvider) ProviderSources() []aws.CredentialSource {
	if s.Source == nil {
		return []aws.CredentialSource{aws.CredentialSourceCode} // If no source has been set, assume this is used directly which means hardcoded creds
	}
	return s.Source
}

// NewStaticCredentialsProvider return a StaticCredentialsProvider initialized with the AWS
// credentials passed in.
func NewStaticCredentialsProvider(key, secret, session string) StaticCredentialsProvider {
	return StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     key,
			SecretAccessKey: secret,
			SessionToken:    session,
		},
	}
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s StaticCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	v := s.Value
	if v.AccessKeyID == "" || v.SecretAccessKey == "" {
		return aws.Credentials{
			Source: StaticCredentialsName,
		}, &StaticCredentialsEmptyError{}
	}

	if len(v.Source) == 0 {
		v.Source = StaticCredentialsName
	}

	return v, nil
}
