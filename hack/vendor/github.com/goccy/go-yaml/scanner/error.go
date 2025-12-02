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

package scanner

import "github.com/goccy/go-yaml/token"

type InvalidTokenError struct {
	Token *token.Token
}

func (e *InvalidTokenError) Error() string {
	return e.Token.Error
}

func ErrInvalidToken(tk *token.Token) *InvalidTokenError {
	return &InvalidTokenError{
		Token: tk,
	}
}
