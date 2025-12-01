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

package waiter

import (
	"context"
	"fmt"

	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
)

// Logger is the Logger middleware used by the waiter to log an attempt
type Logger struct {
	// Attempt is the current attempt to be logged
	Attempt int64
}

// ID representing the Logger middleware
func (*Logger) ID() string {
	return "WaiterLogger"
}

// HandleInitialize performs handling of request in initialize stack step
func (m *Logger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	logger := middleware.GetLogger(ctx)

	logger.Logf(logging.Debug, fmt.Sprintf("attempting waiter request, attempt count: %d", m.Attempt))

	return next.HandleInitialize(ctx, in)
}

// AddLogger is a helper util to add waiter logger after `SetLogger` middleware in
func (m Logger) AddLogger(stack *middleware.Stack) error {
	return stack.Initialize.Insert(&m, "SetLogger", middleware.After)
}
