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

package container

import (
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// NetworkSettings exposes the network settings in the api
type NetworkSettings struct {
	NetworkSettingsBase
	DefaultNetworkSettings
	Networks map[string]*network.EndpointSettings
}

// NetworkSettingsBase holds networking state for a container when inspecting it.
//
// Deprecated: Most fields in NetworkSettingsBase are deprecated. Fields which aren't deprecated will move to
// NetworkSettings in v29.0, and this struct will be removed.
type NetworkSettingsBase struct {
	Bridge     string      // Deprecated: This field is only set when the daemon is started with the --bridge flag specified.
	SandboxID  string      // SandboxID uniquely represents a container's network stack
	SandboxKey string      // SandboxKey identifies the sandbox
	Ports      nat.PortMap // Ports is a collection of PortBinding indexed by Port

	// HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
	//
	// Deprecated: This field is never set and will be removed in a future release.
	HairpinMode bool
	// LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
	//
	// Deprecated: This field is never set and will be removed in a future release.
	LinkLocalIPv6Address string
	// LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
	//
	// Deprecated: This field is never set and will be removed in a future release.
	LinkLocalIPv6PrefixLen int
	SecondaryIPAddresses   []network.Address // Deprecated: This field is never set and will be removed in a future release.
	SecondaryIPv6Addresses []network.Address // Deprecated: This field is never set and will be removed in a future release.
}

// DefaultNetworkSettings holds the networking state for the default bridge, if the container is connected to that
// network.
//
// Deprecated: this struct is deprecated since Docker v1.11 and will be removed in v29. You should look for the default
// network in NetworkSettings.Networks instead.
type DefaultNetworkSettings struct {
	// EndpointID uniquely represents a service endpoint in a Sandbox
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	EndpointID string
	// Gateway holds the gateway address for the network
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	Gateway string
	// GlobalIPv6Address holds network's global IPv6 address
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	GlobalIPv6Address string
	// GlobalIPv6PrefixLen represents mask length of network's global IPv6 address
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	GlobalIPv6PrefixLen int
	// IPAddress holds the IPv4 address for the network
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	IPAddress string
	// IPPrefixLen represents mask length of network's IPv4 address
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	IPPrefixLen int
	// IPv6Gateway holds gateway address specific for IPv6
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	IPv6Gateway string
	// MacAddress holds the MAC address for the network
	//
	// Deprecated: This field will be removed in v29. You should look for the default network in NetworkSettings.Networks instead.
	MacAddress string
}

// NetworkSettingsSummary provides a summary of container's networks
// in /containers/json
type NetworkSettingsSummary struct {
	Networks map[string]*network.EndpointSettings
}
