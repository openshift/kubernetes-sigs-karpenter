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
	"encoding/base64"
	"fmt"
	"io"
)

type base64Encoder struct {
	encoding base64.Encoding
}

func NewBase64Encoder() Encoder {
	return &base64Encoder{encoding: *base64.StdEncoding}
}

func (e *base64Encoder) CanHandleAliases() bool {
	return false
}

func (e *base64Encoder) PrintDocumentSeparator(_ io.Writer) error {
	return nil
}

func (e *base64Encoder) PrintLeadingContent(_ io.Writer, _ string) error {
	return nil
}

func (e *base64Encoder) Encode(writer io.Writer, node *CandidateNode) error {
	if node.guessTagFromCustomType() != "!!str" {
		return fmt.Errorf("cannot encode %v as base64, can only operate on strings", node.Tag)
	}
	_, err := writer.Write([]byte(e.encoding.EncodeToString([]byte(node.Value))))
	return err
}
