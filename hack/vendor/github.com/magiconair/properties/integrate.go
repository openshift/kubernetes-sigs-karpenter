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

// Copyright 2013-2022 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package properties

import "flag"

// MustFlag sets flags that are skipped by dst.Parse when p contains
// the respective key for flag.Flag.Name.
//
// It's use is recommended with command line arguments as in:
//
//	flag.Parse()
//	p.MustFlag(flag.CommandLine)
func (p *Properties) MustFlag(dst *flag.FlagSet) {
	m := make(map[string]*flag.Flag)
	dst.VisitAll(func(f *flag.Flag) {
		m[f.Name] = f
	})
	dst.Visit(func(f *flag.Flag) {
		delete(m, f.Name) // overridden
	})

	for name, f := range m {
		v, ok := p.Get(name)
		if !ok {
			continue
		}

		if err := f.Value.Set(v); err != nil {
			ErrorHandler(err)
		}
	}
}
