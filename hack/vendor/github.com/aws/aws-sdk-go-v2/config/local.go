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

package config

import (
	"fmt"
	"net"
	"net/url"
)

var lookupHostFn = net.LookupHost

func isLoopbackHost(host string) (bool, error) {
	ip := net.ParseIP(host)
	if ip != nil {
		return ip.IsLoopback(), nil
	}

	// Host is not an ip, perform lookup
	addrs, err := lookupHostFn(host)
	if err != nil {
		return false, err
	}
	if len(addrs) == 0 {
		return false, fmt.Errorf("no addrs found for host, %s", host)
	}

	for _, addr := range addrs {
		if !net.ParseIP(addr).IsLoopback() {
			return false, nil
		}
	}

	return true, nil
}

func validateLocalURL(v string) error {
	u, err := url.Parse(v)
	if err != nil {
		return err
	}

	host := u.Hostname()
	if len(host) == 0 {
		return fmt.Errorf("unable to parse host from local HTTP cred provider URL")
	} else if isLoopback, err := isLoopbackHost(host); err != nil {
		return fmt.Errorf("failed to resolve host %q, %v", host, err)
	} else if !isLoopback {
		return fmt.Errorf("invalid endpoint host, %q, only host resolving to loopback addresses are allowed", host)
	}

	return nil
}
