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
	"net/url"
)

type uriEncoder struct {
}

func NewUriEncoder() Encoder {
	return &uriEncoder{}
}

func (e *uriEncoder) CanHandleAliases() bool {
	return false
}

func (e *uriEncoder) PrintDocumentSeparator(_ io.Writer) error {
	return nil
}

func (e *uriEncoder) PrintLeadingContent(_ io.Writer, _ string) error {
	return nil
}

func (e *uriEncoder) Encode(writer io.Writer, node *CandidateNode) error {
	if node.guessTagFromCustomType() != "!!str" {
		return fmt.Errorf("cannot encode %v as URI, can only operate on strings. Please first pipe through another encoding operator to convert the value to a string", node.Tag)
	}
	_, err := writer.Write([]byte(url.QueryEscape(node.Value)))
	return err
}
