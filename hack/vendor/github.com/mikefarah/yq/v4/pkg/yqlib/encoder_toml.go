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

package yqlib

import (
	"fmt"
	"io"
)

type tomlEncoder struct {
}

func NewTomlEncoder() Encoder {
	return &tomlEncoder{}
}

func (te *tomlEncoder) Encode(writer io.Writer, node *CandidateNode) error {
	if node.Kind == ScalarNode {
		return writeString(writer, node.Value+"\n")
	}
	return fmt.Errorf("only scalars (e.g. strings, numbers, booleans) are supported for TOML output at the moment. Please use yaml output format (-oy) until the encoder has been fully implemented")
}

func (te *tomlEncoder) PrintDocumentSeparator(_ io.Writer) error {
	return nil
}

func (te *tomlEncoder) PrintLeadingContent(_ io.Writer, _ string) error {
	return nil
}

func (te *tomlEncoder) CanHandleAliases() bool {
	return false
}
