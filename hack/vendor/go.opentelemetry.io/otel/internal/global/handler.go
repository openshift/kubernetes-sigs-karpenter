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

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package global provides the OpenTelemetry global API.
package global // import "go.opentelemetry.io/otel/internal/global"

import (
	"log"
	"sync/atomic"
)

// ErrorHandler handles irremediable events.
type ErrorHandler interface {
	// Handle handles any error deemed irremediable by an OpenTelemetry
	// component.
	Handle(error)
}

type ErrDelegator struct {
	delegate atomic.Pointer[ErrorHandler]
}

// Compile-time check that delegator implements ErrorHandler.
var _ ErrorHandler = (*ErrDelegator)(nil)

func (d *ErrDelegator) Handle(err error) {
	if eh := d.delegate.Load(); eh != nil {
		(*eh).Handle(err)
		return
	}
	log.Print(err)
}

// setDelegate sets the ErrorHandler delegate.
func (d *ErrDelegator) setDelegate(eh ErrorHandler) {
	d.delegate.Store(&eh)
}
