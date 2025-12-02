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
	"io"
	"sort"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/formatters/svg"
)

var (
	// NoOp formatter.
	NoOp = Register("noop", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, iterator chroma.Iterator) error {
		for t := iterator(); t != chroma.EOF; t = iterator() {
			if _, err := io.WriteString(w, t.Value); err != nil {
				return err
			}
		}
		return nil
	}))
	// Default HTML formatter outputs self-contained HTML.
	htmlFull = Register("html", html.New(html.Standalone(true), html.WithClasses(true))) // nolint
	SVG      = Register("svg", svg.New(svg.EmbedFont("Liberation Mono", svg.FontLiberationMono, svg.WOFF)))
)

// Fallback formatter.
var Fallback = NoOp

// Registry of Formatters.
var Registry = map[string]chroma.Formatter{}

// Names of registered formatters.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get formatter by name.
//
// If the given formatter is not found, the Fallback formatter will be returned.
func Get(name string) chroma.Formatter {
	if f, ok := Registry[name]; ok {
		return f
	}
	return Fallback
}

// Register a named formatter.
func Register(name string, formatter chroma.Formatter) chroma.Formatter {
	Registry[name] = formatter
	return formatter
}
