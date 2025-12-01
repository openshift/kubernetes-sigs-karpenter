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

package client

import (
	"io"
	"net/http"
	"sync/atomic"
)

// KeepAliveTransport drains the remaining body from a response
// so that go will reuse the TCP connections.
// This is not enabled by default because there are servers where
// the response never gets closed and that would make the code hang forever.
// So instead it's provided as a http client middleware that can be used to override
// any request.
func KeepAliveTransport(rt http.RoundTripper) http.RoundTripper {
	return &keepAliveTransport{wrapped: rt}
}

type keepAliveTransport struct {
	wrapped http.RoundTripper
}

func (k *keepAliveTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := k.wrapped.RoundTrip(r)
	if err != nil {
		return resp, err
	}
	resp.Body = &drainingReadCloser{rdr: resp.Body}
	return resp, nil
}

type drainingReadCloser struct {
	rdr     io.ReadCloser
	seenEOF uint32
}

func (d *drainingReadCloser) Read(p []byte) (n int, err error) {
	n, err = d.rdr.Read(p)
	if err == io.EOF || n == 0 {
		atomic.StoreUint32(&d.seenEOF, 1)
	}
	return
}

func (d *drainingReadCloser) Close() error {
	// drain buffer
	if atomic.LoadUint32(&d.seenEOF) != 1 {
		// If the reader side (a HTTP server) is misbehaving, it still may send
		// some bytes, but the closer ignores them to keep the underling
		// connection open.
		_, _ = io.Copy(io.Discard, d.rdr)
	}
	return d.rdr.Close()
}
