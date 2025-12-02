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
	"strings"
)

// Humanize returns first letter of sentence capitalized.
// Common acronyms are capitalized as well.
// Other capital letters in string are left as provided.
//
//	employee_salary = Employee salary
//	employee_id = employee ID
//	employee_mobile_number = Employee mobile number
//	first_Name = First Name
//	firstName = First Name
func Humanize(s string) string {
	return New(s).Humanize().String()
}

// Humanize First letter of sentence capitalized
func (i Ident) Humanize() Ident {
	if len(i.Original) == 0 {
		return New("")
	}

	if strings.TrimSpace(i.Original) == "" {
		return i
	}

	parts := xappend([]string{}, Titleize(i.Parts[0]))
	if len(i.Parts) > 1 {
		parts = xappend(parts, i.Parts[1:]...)
	}

	return New(strings.Join(parts, " "))
}
