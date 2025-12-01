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

//go:build appengine || (!darwin && !dragonfly && !freebsd && !openbsd && !linux && !netbsd && !solaris && !windows)

package fsnotify

import "errors"

type other struct {
	Events chan Event
	Errors chan error
}

var defaultBufferSize = 0

func newBackend(ev chan Event, errs chan error) (backend, error) {
	return nil, errors.New("fsnotify not supported on the current platform")
}
func (w *other) Close() error                              { return nil }
func (w *other) WatchList() []string                       { return nil }
func (w *other) Add(name string) error                     { return nil }
func (w *other) AddWith(name string, opts ...addOpt) error { return nil }
func (w *other) Remove(name string) error                  { return nil }
func (w *other) xSupports(op Op) bool                      { return false }
