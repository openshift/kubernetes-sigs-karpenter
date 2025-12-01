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

// Copyright 2015 Tim Heckman. All rights reserved.
// Copyright 2018-2025 The Gofrs. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

//go:build (!unix && !windows) || plan9

package flock

import (
	"errors"
	"io/fs"
)

func (f *Flock) Lock() error {
	return &fs.PathError{
		Op:   "Lock",
		Path: f.Path(),
		Err:  errors.ErrUnsupported,
	}
}

func (f *Flock) RLock() error {
	return &fs.PathError{
		Op:   "RLock",
		Path: f.Path(),
		Err:  errors.ErrUnsupported,
	}
}

func (f *Flock) Unlock() error {
	return &fs.PathError{
		Op:   "Unlock",
		Path: f.Path(),
		Err:  errors.ErrUnsupported,
	}
}

func (f *Flock) TryLock() (bool, error) {
	return false, f.Lock()
}

func (f *Flock) TryRLock() (bool, error) {
	return false, f.RLock()
}
