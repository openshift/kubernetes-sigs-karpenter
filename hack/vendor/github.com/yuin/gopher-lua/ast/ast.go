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

package ast

type PositionHolder interface {
	Line() int
	SetLine(int)
	LastLine() int
	SetLastLine(int)
}

type Node struct {
	line     int
	lastline int
}

func (self *Node) Line() int {
	return self.line
}

func (self *Node) SetLine(line int) {
	self.line = line
}

func (self *Node) LastLine() int {
	return self.lastline
}

func (self *Node) SetLastLine(line int) {
	self.lastline = line
}
