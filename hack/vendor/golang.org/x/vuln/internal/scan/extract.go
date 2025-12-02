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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"golang.org/x/vuln/internal/derrors"
	"golang.org/x/vuln/internal/vulncheck"
)

const (
	// extractModeID is the unique name of the extract mode protocol
	extractModeID      = "govulncheck-extract"
	extractModeVersion = "0.1.0"
)

// header information for the blob output.
type header struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// runExtract dumps the extracted abstraction of binary at cfg.patterns to out.
// It prints out exactly two blob messages, one with the header and one with
// the vulncheck.Bin as the body.
func runExtract(cfg *config, out io.Writer) (err error) {
	defer derrors.Wrap(&err, "govulncheck")

	bin, err := createBin(cfg.patterns[0])
	if err != nil {
		return err
	}
	sortBin(bin) // sort for easier testing and validation
	header := header{
		Name:    extractModeID,
		Version: extractModeVersion,
	}

	enc := json.NewEncoder(out)

	if err := enc.Encode(header); err != nil {
		return fmt.Errorf("marshaling blob header: %v", err)
	}
	if err := enc.Encode(bin); err != nil {
		return fmt.Errorf("marshaling blob body: %v", err)
	}
	return nil
}

func sortBin(bin *vulncheck.Bin) {
	sort.SliceStable(bin.PkgSymbols, func(i, j int) bool {
		return bin.PkgSymbols[i].Pkg+"."+bin.PkgSymbols[i].Name < bin.PkgSymbols[j].Pkg+"."+bin.PkgSymbols[j].Name
	})
}
