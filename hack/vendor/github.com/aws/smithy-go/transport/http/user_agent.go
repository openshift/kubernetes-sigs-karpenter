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

package http

import (
	"strings"
)

// UserAgentBuilder is a builder for a HTTP User-Agent string.
type UserAgentBuilder struct {
	sb strings.Builder
}

// NewUserAgentBuilder returns a new UserAgentBuilder.
func NewUserAgentBuilder() *UserAgentBuilder {
	return &UserAgentBuilder{sb: strings.Builder{}}
}

// AddKey adds the named component/product to the agent string
func (u *UserAgentBuilder) AddKey(key string) {
	u.appendTo(key)
}

// AddKeyValue adds the named key to the agent string with the given value.
func (u *UserAgentBuilder) AddKeyValue(key, value string) {
	u.appendTo(key + "/" + value)
}

// Build returns the constructed User-Agent string. May be called multiple times.
func (u *UserAgentBuilder) Build() string {
	return u.sb.String()
}

func (u *UserAgentBuilder) appendTo(value string) {
	if u.sb.Len() > 0 {
		u.sb.WriteRune(' ')
	}
	u.sb.WriteString(value)
}
