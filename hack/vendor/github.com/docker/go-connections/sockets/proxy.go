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

package sockets

import (
	"net"
	"os"
	"strings"
)

// GetProxyEnv allows access to the uppercase and the lowercase forms of
// proxy-related variables.  See the Go specification for details on these
// variables. https://golang.org/pkg/net/http/
func GetProxyEnv(key string) string {
	proxyValue := os.Getenv(strings.ToUpper(key))
	if proxyValue == "" {
		return os.Getenv(strings.ToLower(key))
	}
	return proxyValue
}

// DialerFromEnvironment was previously used to configure a net.Dialer to route
// connections through a SOCKS proxy.
// DEPRECATED: SOCKS proxies are now supported by configuring only
// http.Transport.Proxy, and no longer require changing http.Transport.Dial.
// Therefore, only sockets.ConfigureTransport() needs to be called, and any
// sockets.DialerFromEnvironment() calls can be dropped.
func DialerFromEnvironment(direct *net.Dialer) (*net.Dialer, error) {
	return direct, nil
}
