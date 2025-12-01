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

//go:build linux && arm64

package cpu

import (
	"errors"
	"io"
	"os"
	"strings"
)

func readLinuxProcCPUInfo() error {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return err
	}
	defer f.Close()

	var buf [1 << 10]byte // enough for first CPU
	n, err := io.ReadFull(f, buf[:])
	if err != nil && err != io.ErrUnexpectedEOF {
		return err
	}
	in := string(buf[:n])
	const features = "\nFeatures	: "
	i := strings.Index(in, features)
	if i == -1 {
		return errors.New("no CPU features found")
	}
	in = in[i+len(features):]
	if i := strings.Index(in, "\n"); i != -1 {
		in = in[:i]
	}
	m := map[string]*bool{}

	initOptions() // need it early here; it's harmless to call twice
	for _, o := range options {
		m[o.Name] = o.Feature
	}
	// The EVTSTRM field has alias "evstrm" in Go, but Linux calls it "evtstrm".
	m["evtstrm"] = &ARM64.HasEVTSTRM

	for _, f := range strings.Fields(in) {
		if p, ok := m[f]; ok {
			*p = true
		}
	}
	return nil
}
