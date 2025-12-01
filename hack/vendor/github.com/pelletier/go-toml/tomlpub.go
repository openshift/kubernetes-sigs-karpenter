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

package toml

// PubTOMLValue wrapping tomlValue in order to access all properties from outside.
type PubTOMLValue = tomlValue

func (ptv *PubTOMLValue) Value() interface{} {
	return ptv.value
}
func (ptv *PubTOMLValue) Comment() string {
	return ptv.comment
}
func (ptv *PubTOMLValue) Commented() bool {
	return ptv.commented
}
func (ptv *PubTOMLValue) Multiline() bool {
	return ptv.multiline
}
func (ptv *PubTOMLValue) Position() Position {
	return ptv.position
}

func (ptv *PubTOMLValue) SetValue(v interface{}) {
	ptv.value = v
}
func (ptv *PubTOMLValue) SetComment(s string) {
	ptv.comment = s
}
func (ptv *PubTOMLValue) SetCommented(c bool) {
	ptv.commented = c
}
func (ptv *PubTOMLValue) SetMultiline(m bool) {
	ptv.multiline = m
}
func (ptv *PubTOMLValue) SetPosition(p Position) {
	ptv.position = p
}

// PubTree wrapping Tree in order to access all properties from outside.
type PubTree = Tree

func (pt *PubTree) Values() map[string]interface{} {
	return pt.values
}

func (pt *PubTree) Comment() string {
	return pt.comment
}

func (pt *PubTree) Commented() bool {
	return pt.commented
}

func (pt *PubTree) Inline() bool {
	return pt.inline
}

func (pt *PubTree) SetValues(v map[string]interface{}) {
	pt.values = v
}

func (pt *PubTree) SetComment(c string) {
	pt.comment = c
}

func (pt *PubTree) SetCommented(c bool) {
	pt.commented = c
}

func (pt *PubTree) SetInline(i bool) {
	pt.inline = i
}
