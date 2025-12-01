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

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"crypto/tls"
	"errors"
	"net"
)

const nextProtoUnencryptedHTTP2 = "unencrypted_http2"

// unencryptedNetConnFromTLSConn retrieves a net.Conn wrapped in a *tls.Conn.
//
// TLSNextProto functions accept a *tls.Conn.
//
// When passing an unencrypted HTTP/2 connection to a TLSNextProto function,
// we pass a *tls.Conn with an underlying net.Conn containing the unencrypted connection.
// To be extra careful about mistakes (accidentally dropping TLS encryption in a place
// where we want it), the tls.Conn contains a net.Conn with an UnencryptedNetConn method
// that returns the actual connection we want to use.
func unencryptedNetConnFromTLSConn(tc *tls.Conn) (net.Conn, error) {
	conner, ok := tc.NetConn().(interface {
		UnencryptedNetConn() net.Conn
	})
	if !ok {
		return nil, errors.New("http2: TLS conn unexpectedly found in unencrypted handoff")
	}
	return conner.UnencryptedNetConn(), nil
}
