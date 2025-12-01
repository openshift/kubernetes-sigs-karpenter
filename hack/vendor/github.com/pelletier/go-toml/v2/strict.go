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

import (
	"github.com/pelletier/go-toml/v2/internal/danger"
	"github.com/pelletier/go-toml/v2/internal/tracker"
	"github.com/pelletier/go-toml/v2/unstable"
)

type strict struct {
	Enabled bool

	// Tracks the current key being processed.
	key tracker.KeyTracker

	missing []unstable.ParserError
}

func (s *strict) EnterTable(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.key.UpdateTable(node)
}

func (s *strict) EnterArrayTable(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.key.UpdateArrayTable(node)
}

func (s *strict) EnterKeyValue(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.key.Push(node)
}

func (s *strict) ExitKeyValue(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.key.Pop(node)
}

func (s *strict) MissingTable(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.missing = append(s.missing, unstable.ParserError{
		Highlight: keyLocation(node),
		Message:   "missing table",
		Key:       s.key.Key(),
	})
}

func (s *strict) MissingField(node *unstable.Node) {
	if !s.Enabled {
		return
	}

	s.missing = append(s.missing, unstable.ParserError{
		Highlight: keyLocation(node),
		Message:   "missing field",
		Key:       s.key.Key(),
	})
}

func (s *strict) Error(doc []byte) error {
	if !s.Enabled || len(s.missing) == 0 {
		return nil
	}

	err := &StrictMissingError{
		Errors: make([]DecodeError, 0, len(s.missing)),
	}

	for _, derr := range s.missing {
		derr := derr
		err.Errors = append(err.Errors, *wrapDecodeError(doc, &derr))
	}

	return err
}

func keyLocation(node *unstable.Node) []byte {
	k := node.Key()

	hasOne := k.Next()
	if !hasOne {
		panic("should not be called with empty key")
	}

	start := k.Node().Data
	end := k.Node().Data

	for k.Next() {
		end = k.Node().Data
	}

	return danger.BytesRange(start, end)
}
