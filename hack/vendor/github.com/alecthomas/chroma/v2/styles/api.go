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

package styles

import (
	"embed"
	"io/fs"
	"sort"

	"github.com/alecthomas/chroma/v2"
)

//go:embed *.xml
var embedded embed.FS

// Registry of Styles.
var Registry = func() map[string]*chroma.Style {
	registry := map[string]*chroma.Style{}
	// Register all embedded styles.
	files, err := fs.ReadDir(embedded, ".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		r, err := embedded.Open(file.Name())
		if err != nil {
			panic(err)
		}
		style, err := chroma.NewXMLStyle(r)
		if err != nil {
			panic(err)
		}
		registry[style.Name] = style
		_ = r.Close()
	}
	return registry
}()

// Fallback style. Reassign to change the default fallback style.
var Fallback = Registry["swapoff"]

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	Registry[style.Name] = style
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get named style, or Fallback.
func Get(name string) *chroma.Style {
	if style, ok := Registry[name]; ok {
		return style
	}
	return Fallback
}
