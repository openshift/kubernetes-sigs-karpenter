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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package govulncheck

import (
	"encoding/json"

	"io"

	"golang.org/x/vuln/internal/osv"
)

type jsonHandler struct {
	enc *json.Encoder
}

// NewJSONHandler returns a handler that writes govulncheck output as json.
func NewJSONHandler(w io.Writer) Handler {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return &jsonHandler{enc: enc}
}

// Config writes config block in JSON to the underlying writer.
func (h *jsonHandler) Config(config *Config) error {
	return h.enc.Encode(Message{Config: config})
}

// Progress writes a progress message in JSON to the underlying writer.
func (h *jsonHandler) Progress(progress *Progress) error {
	return h.enc.Encode(Message{Progress: progress})
}

// SBOM writes the SBOM block in JSON to the underlying writer.
func (h *jsonHandler) SBOM(sbom *SBOM) error {
	return h.enc.Encode(Message{SBOM: sbom})
}

// OSV writes an osv entry in JSON to the underlying writer.
func (h *jsonHandler) OSV(entry *osv.Entry) error {
	return h.enc.Encode(Message{OSV: entry})
}

// Finding writes a finding in JSON to the underlying writer.
func (h *jsonHandler) Finding(finding *Finding) error {
	return h.enc.Encode(Message{Finding: finding})
}
