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
	"encoding/json"
	"fmt"
	"io"

	"github.com/alecthomas/chroma/v2"
)

// JSON formatter outputs the raw token structures as JSON.
var JSON = Register("json", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, it chroma.Iterator) error {
	if _, err := fmt.Fprintln(w, "["); err != nil {
		return err
	}
	i := 0
	for t := it(); t != chroma.EOF; t = it() {
		if i > 0 {
			if _, err := fmt.Fprintln(w, ","); err != nil {
				return err
			}
		}
		i++
		bytes, err := json.Marshal(t)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprint(w, "  "+string(bytes)); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "]"); err != nil {
		return err
	}
	return nil
}))
