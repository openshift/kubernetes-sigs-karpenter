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

package ext

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// LogError sets the error=true tag on the Span and logs err as an "error" event.
func LogError(span opentracing.Span, err error, fields ...log.Field) {
	Error.Set(span, true)
	ef := []log.Field{
		log.Event("error"),
		log.Error(err),
	}
	ef = append(ef, fields...)
	span.LogFields(ef...)
}
