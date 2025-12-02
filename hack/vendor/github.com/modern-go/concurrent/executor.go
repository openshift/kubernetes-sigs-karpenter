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

package concurrent

import "context"

// Executor replace go keyword to start a new goroutine
// the goroutine should cancel itself if the context passed in has been cancelled
// the goroutine started by the executor, is owned by the executor
// we can cancel all executors owned by the executor just by stop the executor itself
// however Executor interface does not Stop method, the one starting and owning executor
// should use the concrete type of executor, instead of this interface.
type Executor interface {
	// Go starts a new goroutine controlled by the context
	Go(handler func(ctx context.Context))
}
