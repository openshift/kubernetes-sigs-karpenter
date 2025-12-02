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

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.9

package catalog

import "golang.org/x/text/internal/catmsg"

// A Message holds a collection of translations for the same phrase that may
// vary based on the values of substitution arguments.
type Message interface {
	catmsg.Message
}

func firstInSequence(m []Message) catmsg.Message {
	a := []catmsg.Message{}
	for _, m := range m {
		a = append(a, m)
	}
	return catmsg.FirstOf(a)
}
