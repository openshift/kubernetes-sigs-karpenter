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

package termenv

// Hyperlink creates a hyperlink using OSC8.
func Hyperlink(link, name string) string {
	return output.Hyperlink(link, name)
}

// Hyperlink creates a hyperlink using OSC8.
func (o *Output) Hyperlink(link, name string) string {
	return OSC + "8;;" + link + ST + name + OSC + "8;;" + ST
}
