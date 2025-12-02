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

package middleware

import "context"

type (
	serviceIDKey     struct{}
	operationNameKey struct{}
)

// WithServiceID adds a service ID to the context, scoped to middleware stack
// values.
//
// This API is called in the client runtime when bootstrapping an operation and
// should not typically be used directly.
func WithServiceID(parent context.Context, id string) context.Context {
	return WithStackValue(parent, serviceIDKey{}, id)
}

// GetServiceID retrieves the service ID from the context. This is typically
// the service shape's name from its Smithy model. Service clients for specific
// systems (e.g. AWS SDK) may use an alternate designated value.
func GetServiceID(ctx context.Context) string {
	id, _ := GetStackValue(ctx, serviceIDKey{}).(string)
	return id
}

// WithOperationName adds the operation name to the context, scoped to
// middleware stack values.
//
// This API is called in the client runtime when bootstrapping an operation and
// should not typically be used directly.
func WithOperationName(parent context.Context, id string) context.Context {
	return WithStackValue(parent, operationNameKey{}, id)
}

// GetOperationName retrieves the operation name from the context. This is
// typically the operation shape's name from its Smithy model.
func GetOperationName(ctx context.Context) string {
	name, _ := GetStackValue(ctx, operationNameKey{}).(string)
	return name
}
