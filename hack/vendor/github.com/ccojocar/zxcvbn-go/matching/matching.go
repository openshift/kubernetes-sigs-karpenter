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

package matching

import (
	"sort"

	"github.com/ccojocar/zxcvbn-go/adjacency"
	"github.com/ccojocar/zxcvbn-go/frequency"
	"github.com/ccojocar/zxcvbn-go/match"
)

var (
	dictionaryMatchers []match.Matcher
	matchers           []match.Matcher
	adjacencyGraphs    []adjacency.Graph
	l33tTable          adjacency.Graph

	sequences map[string]string
)

func init() {
	loadFrequencyList()
}

// Omnimatch runs all matchers against the password
func Omnimatch(password string, userInputs []string, filters ...func(match.Matcher) bool) (matches []match.Match) {
	// Can I run into the issue where nil is not equal to nil?
	if dictionaryMatchers == nil || adjacencyGraphs == nil {
		loadFrequencyList()
	}

	if userInputs != nil {
		userInputMatcher := buildDictMatcher("user_inputs", buildRankedDict(userInputs))
		matches = userInputMatcher(password)
	}

	for _, matcher := range matchers {
		shouldBeFiltered := false
		for i := range filters {
			if filters[i](matcher) {
				shouldBeFiltered = true
				break
			}
		}
		if !shouldBeFiltered {
			matches = append(matches, matcher.MatchingFunc(password)...)
		}
	}
	sort.Sort(match.Matches(matches))
	return matches
}

func loadFrequencyList() {
	for n, list := range frequency.Lists {
		dictionaryMatchers = append(dictionaryMatchers, match.Matcher{MatchingFunc: buildDictMatcher(n, buildRankedDict(list.List)), ID: n})
	}

	l33tTable = adjacency.GraphMap["l33t"]

	adjacencyGraphs = append(adjacencyGraphs, adjacency.GraphMap["qwerty"])
	adjacencyGraphs = append(adjacencyGraphs, adjacency.GraphMap["dvorak"])
	adjacencyGraphs = append(adjacencyGraphs, adjacency.GraphMap["keypad"])
	adjacencyGraphs = append(adjacencyGraphs, adjacency.GraphMap["macKeypad"])

	// l33tFilePath, _ := filepath.Abs("adjacency/L33t.json")
	// L33T_TABLE = adjacency.GetAdjancencyGraphFromFile(l33tFilePath, "l33t")

	sequences = make(map[string]string)
	sequences["lower"] = "abcdefghijklmnopqrstuvwxyz"
	sequences["upper"] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sequences["digits"] = "0123456789"

	matchers = append(matchers, dictionaryMatchers...)
	matchers = append(matchers, match.Matcher{MatchingFunc: spatialMatch, ID: spatialMatcherName})
	matchers = append(matchers, match.Matcher{MatchingFunc: repeatMatch, ID: repeatMatcherName})
	matchers = append(matchers, match.Matcher{MatchingFunc: sequenceMatch, ID: sequenceMatcherName})
	matchers = append(matchers, match.Matcher{MatchingFunc: l33tMatch, ID: L33TMatcherName})
	matchers = append(matchers, match.Matcher{MatchingFunc: dateSepMatcher, ID: dateSepMatcherName})
	matchers = append(matchers, match.Matcher{MatchingFunc: dateWithoutSepMatch, ID: dateWithOutSepMatcherName})
}
