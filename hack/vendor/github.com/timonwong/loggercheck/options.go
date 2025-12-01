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

package loggercheck

import (
	"github.com/timonwong/loggercheck/internal/sets"
)

type Option func(*loggercheck)

func WithDisable(disable []string) Option {
	return func(l *loggercheck) {
		l.disable = sets.NewString(disable...)
	}
}

func WithRules(customRules []string) Option {
	return func(l *loggercheck) {
		l.rules = customRules
	}
}

func WithRequireStringKey(requireStringKey bool) Option {
	return func(l *loggercheck) {
		l.requireStringKey = requireStringKey
	}
}

func WithNoPrintfLike(noPrintfLike bool) Option {
	return func(l *loggercheck) {
		l.noPrintfLike = noPrintfLike
	}
}
