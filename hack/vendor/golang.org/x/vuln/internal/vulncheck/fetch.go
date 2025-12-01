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

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vulncheck

import (
	"context"
	"fmt"

	"golang.org/x/tools/go/packages"
	"golang.org/x/vuln/internal/client"
)

// FetchVulnerabilities fetches vulnerabilities that affect the supplied modules.
func FetchVulnerabilities(ctx context.Context, c *client.Client, modules []*packages.Module) ([]*ModVulns, error) {
	mreqs := make([]*client.ModuleRequest, len(modules))
	for i, mod := range modules {
		modPath := mod.Path
		if mod.Replace != nil {
			modPath = mod.Replace.Path
		}
		mreqs[i] = &client.ModuleRequest{
			Path: modPath,
		}
	}
	resps, err := c.ByModules(ctx, mreqs)
	if err != nil {
		return nil, fmt.Errorf("fetching vulnerabilities: %v", err)
	}
	var mv []*ModVulns
	for i, resp := range resps {
		if len(resp.Entries) == 0 {
			continue
		}
		mv = append(mv, &ModVulns{
			Module: modules[i],
			Vulns:  resp.Entries,
		})
	}
	return mv, nil
}
