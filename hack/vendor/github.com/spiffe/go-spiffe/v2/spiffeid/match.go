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

package spiffeid

import "fmt"

// Matcher is used to match a SPIFFE ID.
type Matcher func(ID) error

// MatchAny matches any SPIFFE ID.
func MatchAny() Matcher {
	return Matcher(func(actual ID) error {
		return nil
	})
}

// MatchID matches a specific SPIFFE ID.
func MatchID(expected ID) Matcher {
	return Matcher(func(actual ID) error {
		if actual != expected {
			return fmt.Errorf("unexpected ID %q", actual)
		}
		return nil
	})
}

// MatchOneOf matches any SPIFFE ID in the given list of IDs.
func MatchOneOf(expected ...ID) Matcher {
	set := make(map[ID]struct{})
	for _, id := range expected {
		set[id] = struct{}{}
	}
	return Matcher(func(actual ID) error {
		if _, ok := set[actual]; !ok {
			return fmt.Errorf("unexpected ID %q", actual)
		}
		return nil
	})
}

// MatchMemberOf matches any SPIFFE ID in the given trust domain.
func MatchMemberOf(expected TrustDomain) Matcher {
	return Matcher(func(actual ID) error {
		if !actual.MemberOf(expected) {
			return fmt.Errorf("unexpected trust domain %q", actual.TrustDomain())
		}
		return nil
	})
}
