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

package yit

import (
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	All = func(node *yaml.Node) bool {
		return true
	}

	None = func(node *yaml.Node) bool {
		return false
	}

	StringValue = Intersect(
		WithKind(yaml.ScalarNode),
		WithShortTag("!!str"),
	)
)

func Intersect(ps ...Predicate) Predicate {
	return func(node *yaml.Node) bool {
		for _, p := range ps {
			if !p(node) {
				return false
			}
		}
		return true
	}
}

func Union(ps ...Predicate) Predicate {
	return func(node *yaml.Node) bool {
		for _, p := range ps {
			if p(node) {
				return true
			}
		}
		return false
	}
}

func Negate(p Predicate) Predicate {
	return func(node *yaml.Node) bool {
		return !p(node)
	}
}

func WithStringValue(value string) Predicate {
	return Intersect(
		StringValue,
		func(node *yaml.Node) bool {
			return node.Value == value
		},
	)
}

func WithShortTag(tag string) Predicate {
	return func(node *yaml.Node) bool {
		return node.ShortTag() == tag
	}
}

func WithValue(value string) Predicate {
	return func(node *yaml.Node) bool {
		return node.Value == value
	}
}

func WithKind(kind yaml.Kind) Predicate {
	return func(node *yaml.Node) bool {
		return node.Kind == kind
	}
}

func WithMapKey(key string) Predicate {
	return func(node *yaml.Node) bool {
		return FromNode(node).MapKeys().AnyMatch(WithValue(key))
	}
}

func WithMapValue(value string) Predicate {
	return func(node *yaml.Node) bool {
		return FromNode(node).MapValues().AnyMatch(WithValue(value))
	}
}

func WithMapKeyValue(keyPredicate, valuePredicate Predicate) Predicate {
	return Intersect(
		WithKind(yaml.MappingNode),
		func(node *yaml.Node) bool {
			for i := 0; i < len(node.Content); i += 2 {
				key := node.Content[i]
				value := node.Content[i+1]
				if keyPredicate(key) && valuePredicate(value) {
					return true
				}
			}
			return false
		},
	)
}

func WithPrefix(prefix string) Predicate {
	return func(node *yaml.Node) bool {
		return strings.HasPrefix(node.Value, prefix)
	}
}

func WithSuffix(suffix string) Predicate {
	return func(node *yaml.Node) bool {
		return strings.HasSuffix(node.Value, suffix)
	}
}
