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

package flect

import (
	"fmt"
	"strconv"
)

// Ordinalize converts a number to an ordinal version
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func Ordinalize(s string) string {
	return New(s).Ordinalize().String()
}

// Ordinalize converts a number to an ordinal version
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func (i Ident) Ordinalize() Ident {
	number, err := strconv.Atoi(i.Original)
	if err != nil {
		return i
	}
	var s string
	switch abs(number) % 100 {
	case 11, 12, 13:
		s = fmt.Sprintf("%dth", number)
	default:
		switch abs(number) % 10 {
		case 1:
			s = fmt.Sprintf("%dst", number)
		case 2:
			s = fmt.Sprintf("%dnd", number)
		case 3:
			s = fmt.Sprintf("%drd", number)
		}
	}
	if s != "" {
		return New(s)
	}
	return New(fmt.Sprintf("%dth", number))
}
