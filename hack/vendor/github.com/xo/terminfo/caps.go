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

package terminfo

// BoolCapName returns the bool capability name.
func BoolCapName(i int) string {
	return boolCapNames[2*i]
}

// BoolCapNameShort returns the short bool capability name.
func BoolCapNameShort(i int) string {
	return boolCapNames[2*i+1]
}

// NumCapName returns the num capability name.
func NumCapName(i int) string {
	return numCapNames[2*i]
}

// NumCapNameShort returns the short num capability name.
func NumCapNameShort(i int) string {
	return numCapNames[2*i+1]
}

// StringCapName returns the string capability name.
func StringCapName(i int) string {
	return stringCapNames[2*i]
}

// StringCapNameShort returns the short string capability name.
func StringCapNameShort(i int) string {
	return stringCapNames[2*i+1]
}
