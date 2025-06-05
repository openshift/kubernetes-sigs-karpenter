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

//go:build amd64 && !appengine && !noasm && gc
// +build amd64,!appengine,!noasm,gc

// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.

package zstd

// matchLen returns how many bytes match in a and b
//
// It assumes that:
//
//	len(a) <= len(b) and len(a) > 0
//
//go:noescape
func matchLen(a []byte, b []byte) int
