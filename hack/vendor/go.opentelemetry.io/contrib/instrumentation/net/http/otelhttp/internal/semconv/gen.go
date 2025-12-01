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

package semconv // import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/internal/semconv"

// Generate semconv package:
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/bench_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=bench_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/common_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=common_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/env.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=env.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/env_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=env_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/httpconv.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=httpconv.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/httpconv_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=httpconv_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/httpconvtest_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=httpconvtest_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/util.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=util.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/util_test.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=util_test.go
//go:generate gotmpl --body=../../../../../../internal/shared/semconv/v1.20.0.go.tmpl "--data={ \"pkg\": \"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\" }" --out=v1.20.0.go
