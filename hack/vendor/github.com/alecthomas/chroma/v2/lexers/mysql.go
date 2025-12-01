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

package lexers

import (
	"regexp"
)

var (
	mysqlAnalyserNameBetweenBacktickRe = regexp.MustCompile("`[a-zA-Z_]\\w*`")
	mysqlAnalyserNameBetweenBracketRe  = regexp.MustCompile(`\[[a-zA-Z_]\w*\]`)
)

func init() { // nolint: gochecknoinits
	Get("mysql").
		SetAnalyser(func(text string) float32 {
			nameBetweenBacktickCount := len(mysqlAnalyserNameBetweenBacktickRe.FindAllString(text, -1))
			nameBetweenBracketCount := len(mysqlAnalyserNameBetweenBracketRe.FindAllString(text, -1))

			var result float32

			// Same logic as above in the TSQL analysis.
			dialectNameCount := nameBetweenBacktickCount + nameBetweenBracketCount
			if dialectNameCount >= 1 && nameBetweenBacktickCount >= (2*nameBetweenBracketCount) {
				// Found at least twice as many `name` as [name].
				result += 0.5
			} else if nameBetweenBacktickCount > nameBetweenBracketCount {
				result += 0.2
			} else if nameBetweenBacktickCount > 0 {
				result += 0.1
			}

			return result
		})
}
