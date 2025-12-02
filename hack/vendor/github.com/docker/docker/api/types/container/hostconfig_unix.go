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

//go:build !windows

package container

import "github.com/docker/docker/api/types/network"

// IsValid indicates if an isolation technology is valid
func (i Isolation) IsValid() bool {
	return i.IsDefault()
}

// IsBridge indicates whether container uses the bridge network stack
func (n NetworkMode) IsBridge() bool {
	return n == network.NetworkBridge
}

// IsHost indicates whether container uses the host network stack.
func (n NetworkMode) IsHost() bool {
	return n == network.NetworkHost
}

// IsUserDefined indicates user-created network
func (n NetworkMode) IsUserDefined() bool {
	return !n.IsDefault() && !n.IsBridge() && !n.IsHost() && !n.IsNone() && !n.IsContainer()
}

// NetworkName returns the name of the network stack.
func (n NetworkMode) NetworkName() string {
	switch {
	case n.IsDefault():
		return network.NetworkDefault
	case n.IsBridge():
		return network.NetworkBridge
	case n.IsHost():
		return network.NetworkHost
	case n.IsNone():
		return network.NetworkNone
	case n.IsContainer():
		return "container"
	case n.IsUserDefined():
		return n.UserDefined()
	default:
		return ""
	}
}
