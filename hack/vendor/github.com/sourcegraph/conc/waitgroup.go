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

package conc

import (
	"sync"

	"github.com/sourcegraph/conc/panics"
)

// NewWaitGroup creates a new WaitGroup.
func NewWaitGroup() *WaitGroup {
	return &WaitGroup{}
}

// WaitGroup is the primary building block for scoped concurrency.
// Goroutines can be spawned in the WaitGroup with the Go method,
// and calling Wait() will ensure that each of those goroutines exits
// before continuing. Any panics in a child goroutine will be caught
// and propagated to the caller of Wait().
//
// The zero value of WaitGroup is usable, just like sync.WaitGroup.
// Also like sync.WaitGroup, it must not be copied after first use.
type WaitGroup struct {
	wg sync.WaitGroup
	pc panics.Catcher
}

// Go spawns a new goroutine in the WaitGroup.
func (h *WaitGroup) Go(f func()) {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		h.pc.Try(f)
	}()
}

// Wait will block until all goroutines spawned with Go exit and will
// propagate any panics spawned in a child goroutine.
func (h *WaitGroup) Wait() {
	h.wg.Wait()

	// Propagate a panic if we caught one from a child goroutine.
	h.pc.Repanic()
}

// WaitAndRecover will block until all goroutines spawned with Go exit and
// will return a *panics.Recovered if one of the child goroutines panics.
func (h *WaitGroup) WaitAndRecover() *panics.Recovered {
	h.wg.Wait()

	// Return a recovered panic if we caught one from a child goroutine.
	return h.pc.Recovered()
}
