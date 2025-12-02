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

// Copyright 2012 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transport contains HTTP transports used to make
// authenticated API requests.
//
// This package is DEPRECATED. Users should instead use,
//
//	service, err := NewService(..., option.WithAPIKey(...))
package transport

import (
	"errors"
	"net/http"
)

// APIKey is an HTTP Transport which wraps an underlying transport and
// appends an API Key "key" parameter to the URL of outgoing requests.
//
// Deprecated: please use NewService(..., option.WithAPIKey(...)) instead.
type APIKey struct {
	// Key is the API Key to set on requests.
	Key string

	// Transport is the underlying HTTP transport.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper
}

func (t *APIKey) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
		if rt == nil {
			return nil, errors.New("googleapi/transport: no Transport specified or available")
		}
	}
	newReq := *req
	args := newReq.URL.Query()
	args.Set("key", t.Key)
	newReq.URL.RawQuery = args.Encode()
	return rt.RoundTrip(&newReq)
}
