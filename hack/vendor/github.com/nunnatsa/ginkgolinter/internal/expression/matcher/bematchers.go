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

package matcher

import "github.com/nunnatsa/ginkgolinter/internal/expression/value"

type BeIdenticalToMatcher struct {
	value.Value
}

func (BeIdenticalToMatcher) Type() Type {
	return BeIdenticalToMatcherType
}

func (BeIdenticalToMatcher) MatcherName() string {
	return beIdenticalTo
}

type BeEquivalentToMatcher struct {
	value.Value
}

func (BeEquivalentToMatcher) Type() Type {
	return BeEquivalentToMatcherType
}

func (BeEquivalentToMatcher) MatcherName() string {
	return beEquivalentTo
}

type BeZeroMatcher struct{}

func (BeZeroMatcher) Type() Type {
	return BeZeroMatcherType
}

func (BeZeroMatcher) MatcherName() string {
	return beZero
}

type BeEmptyMatcher struct{}

func (BeEmptyMatcher) Type() Type {
	return BeEmptyMatcherType
}

func (BeEmptyMatcher) MatcherName() string {
	return beEmpty
}

type BeTrueMatcher struct{}

func (BeTrueMatcher) Type() Type {
	return BeTrueMatcherType | BoolValueTrue
}

func (BeTrueMatcher) MatcherName() string {
	return beTrue
}

type BeFalseMatcher struct{}

func (BeFalseMatcher) Type() Type {
	return BeFalseMatcherType | BoolValueFalse
}

func (BeFalseMatcher) MatcherName() string {
	return beFalse
}

type BeNilMatcher struct{}

func (BeNilMatcher) Type() Type {
	return BeNilMatcherType
}

func (BeNilMatcher) MatcherName() string {
	return beNil
}
