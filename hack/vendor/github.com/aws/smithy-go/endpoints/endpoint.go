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

package transport

import (
	"net/http"
	"net/url"

	"github.com/aws/smithy-go"
)

// Endpoint is the endpoint object returned by Endpoint resolution V2
type Endpoint struct {
	// The complete URL minimally specfiying the scheme and host.
	// May optionally specify the port and base path component.
	URI url.URL

	// An optional set of headers to be sent using transport layer headers.
	Headers http.Header

	// A grab-bag property map of endpoint attributes. The
	// values present here are subject to change, or being add/removed at any
	// time.
	Properties smithy.Properties
}
