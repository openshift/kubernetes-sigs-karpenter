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

package spec

import (
	"embed"
	"path"
)

//go:embed schemas/*.json schemas/*/*.json
var assets embed.FS

func jsonschemaDraft04JSONBytes() ([]byte, error) {
	return assets.ReadFile(path.Join("schemas", "jsonschema-draft-04.json"))
}

func v2SchemaJSONBytes() ([]byte, error) {
	return assets.ReadFile(path.Join("schemas", "v2", "schema.json"))
}
