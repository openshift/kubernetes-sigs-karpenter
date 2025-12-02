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

package protogetter

import (
	"go/token"
)

type PosFilter struct {
	positions       map[token.Pos]struct{}
	alreadyReplaced map[string]map[int][2]int // map[filename][line][start, end]
}

func NewPosFilter() *PosFilter {
	return &PosFilter{
		positions:       make(map[token.Pos]struct{}),
		alreadyReplaced: make(map[string]map[int][2]int),
	}
}

func (f *PosFilter) IsFiltered(pos token.Pos) bool {
	_, ok := f.positions[pos]
	return ok
}

func (f *PosFilter) AddPos(pos token.Pos) {
	f.positions[pos] = struct{}{}
}

func (f *PosFilter) IsAlreadyReplaced(fset *token.FileSet, pos, end token.Pos) bool {
	filePos := fset.Position(pos)
	fileEnd := fset.Position(end)

	lines, ok := f.alreadyReplaced[filePos.Filename]
	if !ok {
		return false
	}

	lineRange, ok := lines[filePos.Line]
	if !ok {
		return false
	}

	if lineRange[0] <= filePos.Offset && fileEnd.Offset <= lineRange[1] {
		return true
	}

	return false
}

func (f *PosFilter) AddAlreadyReplaced(fset *token.FileSet, pos, end token.Pos) {
	filePos := fset.Position(pos)
	fileEnd := fset.Position(end)

	lines, ok := f.alreadyReplaced[filePos.Filename]
	if !ok {
		lines = make(map[int][2]int)
		f.alreadyReplaced[filePos.Filename] = lines
	}

	lineRange, ok := lines[filePos.Line]
	if ok && lineRange[0] <= filePos.Offset && fileEnd.Offset <= lineRange[1] {
		return
	}

	lines[filePos.Line] = [2]int{filePos.Offset, fileEnd.Offset}
}
