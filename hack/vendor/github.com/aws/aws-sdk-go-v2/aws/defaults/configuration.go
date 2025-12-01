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

package defaults

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Configuration is the set of SDK configuration options that are determined based
// on the configured DefaultsMode.
type Configuration struct {
	// RetryMode is the configuration's default retry mode API clients should
	// use for constructing a Retryer.
	RetryMode aws.RetryMode

	// ConnectTimeout is the maximum amount of time a dial will wait for
	// a connect to complete.
	//
	// See https://pkg.go.dev/net#Dialer.Timeout
	ConnectTimeout *time.Duration

	// TLSNegotiationTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake.
	//
	// See https://pkg.go.dev/net/http#Transport.TLSHandshakeTimeout
	TLSNegotiationTimeout *time.Duration
}

// GetConnectTimeout returns the ConnectTimeout value, returns false if the value is not set.
func (c *Configuration) GetConnectTimeout() (time.Duration, bool) {
	if c.ConnectTimeout == nil {
		return 0, false
	}
	return *c.ConnectTimeout, true
}

// GetTLSNegotiationTimeout returns the TLSNegotiationTimeout value, returns false if the value is not set.
func (c *Configuration) GetTLSNegotiationTimeout() (time.Duration, bool) {
	if c.TLSNegotiationTimeout == nil {
		return 0, false
	}
	return *c.TLSNegotiationTimeout, true
}
