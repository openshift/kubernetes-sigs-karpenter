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

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "go.opentelemetry.io/otel/trace/internal/telemetry"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Scope is the identifying values of the instrumentation scope.
type Scope struct {
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	Attrs        []Attr `json:"attributes,omitempty"`
	DroppedAttrs uint32 `json:"droppedAttributesCount,omitempty"`
}

// UnmarshalJSON decodes the OTLP formatted JSON contained in data into r.
func (s *Scope) UnmarshalJSON(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))

	t, err := decoder.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return errors.New("invalid Scope type")
	}

	for decoder.More() {
		keyIface, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// Empty.
				return nil
			}
			return err
		}

		key, ok := keyIface.(string)
		if !ok {
			return fmt.Errorf("invalid Scope field: %#v", keyIface)
		}

		switch key {
		case "name":
			err = decoder.Decode(&s.Name)
		case "version":
			err = decoder.Decode(&s.Version)
		case "attributes":
			err = decoder.Decode(&s.Attrs)
		case "droppedAttributesCount", "dropped_attributes_count":
			err = decoder.Decode(&s.DroppedAttrs)
		default:
			// Skip unknown.
		}

		if err != nil {
			return err
		}
	}
	return nil
}
