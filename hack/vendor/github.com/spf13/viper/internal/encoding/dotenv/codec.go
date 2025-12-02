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

package dotenv

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/subosito/gotenv"
)

const keyDelimiter = "_"

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for encoding data containing environment variables
// (commonly called as dotenv format).
type Codec struct{}

func (Codec) Encode(v map[string]any) ([]byte, error) {
	flattened := map[string]any{}

	flattened = flattenAndMergeMap(flattened, v, "", keyDelimiter)

	keys := make([]string, 0, len(flattened))

	for key := range flattened {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var buf bytes.Buffer

	for _, key := range keys {
		_, err := buf.WriteString(fmt.Sprintf("%v=%v\n", strings.ToUpper(key), flattened[key]))
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (Codec) Decode(b []byte, v map[string]any) error {
	var buf bytes.Buffer

	_, err := buf.Write(b)
	if err != nil {
		return err
	}

	env, err := gotenv.StrictParse(&buf)
	if err != nil {
		return err
	}

	for key, value := range env {
		v[key] = value
	}

	return nil
}
