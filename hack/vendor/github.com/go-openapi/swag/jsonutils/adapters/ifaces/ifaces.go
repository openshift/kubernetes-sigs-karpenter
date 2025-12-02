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
	_ "encoding/json" // for documentation purpose
	"iter"
)

// Ordered knows how to iterate over the (key,value) pairs of a JSON object.
type Ordered interface {
	OrderedItems() iter.Seq2[string, any]
}

// SetOrdered knows how to append or update the keys of a JSON object,
// given an iterator over (key,value) pairs.
//
// If the provided iterator is nil then the receiver should be set to nil.
type SetOrdered interface {
	SetOrderedItems(iter.Seq2[string, any])
}

// OrderedMap represent a JSON object (i.e. like a map[string,any]),
// and knows how to serialize and deserialize JSON with the order of keys maintained.
type OrderedMap interface {
	Ordered
	SetOrdered

	OrderedMarshalJSON() ([]byte, error)
	OrderedUnmarshalJSON([]byte) error
}

// MarshalAdapter behaves likes the standard library [json.Marshal].
type MarshalAdapter interface {
	Poolable

	Marshal(any) ([]byte, error)
}

// OrderedMarshalAdapter behaves likes the standard library [json.Marshal], preserving the order of keys in objects.
type OrderedMarshalAdapter interface {
	Poolable

	OrderedMarshal(Ordered) ([]byte, error)
}

// UnmarshalAdapter behaves likes the standard library [json.Unmarshal].
type UnmarshalAdapter interface {
	Poolable

	Unmarshal([]byte, any) error
}

// OrderedUnmarshalAdapter behaves likes the standard library [json.Unmarshal], preserving the order of keys in objects.
type OrderedUnmarshalAdapter interface {
	Poolable

	OrderedUnmarshal([]byte, SetOrdered) error
}

// Adapter exposes an interface like the standard [json] library.
type Adapter interface {
	MarshalAdapter
	UnmarshalAdapter

	OrderedAdapter
}

// OrderedAdapter exposes interfaces to process JSON and keep the order of object keys.
type OrderedAdapter interface {
	OrderedMarshalAdapter
	OrderedUnmarshalAdapter
	NewOrderedMap(capacity int) OrderedMap
}

type Poolable interface {
	// Self-redeem: for [Adapter] s that are allocated from a pool.
	// The [Adapter] must not be used after calling [Redeem].
	Redeem()

	// Reset the state of the [Adapter], if any.
	Reset()
}
