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

package ifaces

import (
	"strings"
)

// Capability indicates what a JSON adapter is capable of.
type Capability uint8

const (
	CapabilityMarshalJSON Capability = 1 << iota
	CapabilityUnmarshalJSON
	CapabilityOrderedMarshalJSON
	CapabilityOrderedUnmarshalJSON
	CapabilityOrderedMap
)

func (c Capability) String() string {
	switch c {
	case CapabilityMarshalJSON:
		return "MarshalJSON"
	case CapabilityUnmarshalJSON:
		return "UnmarshalJSON"
	case CapabilityOrderedMarshalJSON:
		return "OrderedMarshalJSON"
	case CapabilityOrderedUnmarshalJSON:
		return "OrderedUnmarshalJSON"
	case CapabilityOrderedMap:
		return "OrderedMap"
	default:
		return "<unknown>"
	}
}

// Capabilities holds several unitary capability flags
type Capabilities uint8

// Has some capability flag enabled.
func (c Capabilities) Has(capability Capability) bool {
	return Capability(c)&capability > 0
}

func (c Capabilities) String() string {
	var w strings.Builder

	first := true
	for _, capability := range []Capability{
		CapabilityMarshalJSON,
		CapabilityUnmarshalJSON,
		CapabilityOrderedMarshalJSON,
		CapabilityOrderedUnmarshalJSON,
		CapabilityOrderedMap,
	} {
		if c.Has(capability) {
			if !first {
				w.WriteByte('|')
			} else {
				first = false
			}
			w.WriteString(capability.String())
		}
	}

	return w.String()
}

const (
	AllCapabilities Capabilities = Capabilities(uint8(CapabilityMarshalJSON) |
		uint8(CapabilityUnmarshalJSON) |
		uint8(CapabilityOrderedMarshalJSON) |
		uint8(CapabilityOrderedUnmarshalJSON) |
		uint8(CapabilityOrderedMap))

	AllUnorderedCapabilities Capabilities = Capabilities(uint8(CapabilityMarshalJSON) | uint8(CapabilityUnmarshalJSON))
)

// RegistryEntry describes how any given adapter registers its capabilities to the [Registrar].
type RegistryEntry struct {
	Who         string
	What        Capabilities
	Constructor func() Adapter
	Support     func(what Capability, value any) bool
}

// Registrar is a type that knows how to keep registration calls from adapters.
type Registrar interface {
	RegisterFor(RegistryEntry)
}
