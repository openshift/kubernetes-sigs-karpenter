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

package formatters

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma/v2"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, it chroma.Iterator) error {
	for t := it(); t != chroma.EOF; t = it() {
		if _, err := fmt.Fprintln(w, t.GoString()); err != nil {
			return err
		}
	}
	return nil
}))
