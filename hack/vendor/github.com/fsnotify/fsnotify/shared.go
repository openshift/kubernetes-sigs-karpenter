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

package fsnotify

import "sync"

type shared struct {
	Events chan Event
	Errors chan error
	done   chan struct{}
	mu     sync.Mutex
}

func newShared(ev chan Event, errs chan error) *shared {
	return &shared{
		Events: ev,
		Errors: errs,
		done:   make(chan struct{}),
	}
}

// Returns true if the event was sent, or false if watcher is closed.
func (w *shared) sendEvent(e Event) bool {
	if e.Op == 0 {
		return true
	}
	select {
	case <-w.done:
		return false
	case w.Events <- e:
		return true
	}
}

// Returns true if the error was sent, or false if watcher is closed.
func (w *shared) sendError(err error) bool {
	if err == nil {
		return true
	}
	select {
	case <-w.done:
		return false
	case w.Errors <- err:
		return true
	}
}

func (w *shared) isClosed() bool {
	select {
	case <-w.done:
		return true
	default:
		return false
	}
}

// Mark as closed; returns true if it was already closed.
func (w *shared) close() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.isClosed() {
		return true
	}
	close(w.done)
	return false
}
