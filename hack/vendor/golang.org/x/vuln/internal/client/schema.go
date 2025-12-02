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

package client

import (
	"encoding/json"
	"path"
	"sort"
	"time"
)

const (
	idDir    = "ID"
	indexDir = "index"
)

var (
	dbEndpoint      = path.Join(indexDir, "db")
	modulesEndpoint = path.Join(indexDir, "modules")
)

func entryEndpoint(id string) string {
	return path.Join(idDir, id)
}

// dbMeta contains metadata about the database itself.
type dbMeta struct {
	// Modified is the time the database was last modified, calculated
	// as the most recent time any single OSV entry was modified.
	Modified time.Time `json:"modified"`
}

// moduleMeta contains metadata about a Go module that has one
// or more vulnerabilities in the database.
//
// Found in the "index/modules" endpoint of the vulnerability database.
type moduleMeta struct {
	// Path is the module path.
	Path string `json:"path"`
	// Vulns is a list of vulnerabilities that affect this module.
	Vulns []moduleVuln `json:"vulns"`
}

// moduleVuln contains metadata about a vulnerability that affects
// a certain module.
type moduleVuln struct {
	// ID is a unique identifier for the vulnerability.
	// The Go vulnerability database issues IDs of the form
	// GO-<YEAR>-<ENTRYID>.
	ID string `json:"id"`
	// Modified is the time the vuln was last modified.
	Modified time.Time `json:"modified"`
	// Fixed is the latest version that introduces a fix for the
	// vulnerability, in SemVer 2.0.0 format, with no leading "v" prefix.
	Fixed string `json:"fixed,omitempty"`
}

// modulesIndex represents an in-memory modules index.
type modulesIndex map[string]*moduleMeta

func (m modulesIndex) MarshalJSON() ([]byte, error) {
	modules := make([]*moduleMeta, 0, len(m))
	for _, module := range m {
		modules = append(modules, module)
	}
	sort.SliceStable(modules, func(i, j int) bool {
		return modules[i].Path < modules[j].Path
	})
	for _, module := range modules {
		sort.SliceStable(module.Vulns, func(i, j int) bool {
			return module.Vulns[i].ID < module.Vulns[j].ID
		})
	}
	return json.Marshal(modules)
}
