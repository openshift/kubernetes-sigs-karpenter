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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/vuln/internal/derrors"
	"golang.org/x/vuln/internal/osv"
)

type source interface {
	// get returns the raw, uncompressed bytes at the
	// requested endpoint, which should be bare with no file extensions
	// (e.g., "index/modules" instead of "index/modules.json.gz").
	// It errors if the endpoint cannot be reached or does not exist
	// in the expected form.
	get(ctx context.Context, endpoint string) ([]byte, error)
}

func newHTTPSource(url string, opts *Options) *httpSource {
	c := http.DefaultClient
	if opts != nil && opts.HTTPClient != nil {
		c = opts.HTTPClient
	}
	return &httpSource{url: url, c: c}
}

// httpSource reads a vulnerability database from an http(s) source.
type httpSource struct {
	url string
	c   *http.Client
}

func (hs *httpSource) get(ctx context.Context, endpoint string) (_ []byte, err error) {
	derrors.Wrap(&err, "get(%s)", endpoint)

	method := http.MethodGet
	reqURL := fmt.Sprintf("%s/%s", hs.url, endpoint+".json.gz")
	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hs.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %s %s returned unexpected status: %s", method, reqURL, resp.Status)
	}

	// Uncompress the result.
	r, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}

func newLocalSource(dir string) *localSource {
	return &localSource{fs: os.DirFS(dir)}
}

// localSource reads a vulnerability database from a local file system.
type localSource struct {
	fs fs.FS
}

func (ls *localSource) get(ctx context.Context, endpoint string) (_ []byte, err error) {
	derrors.Wrap(&err, "get(%s)", endpoint)

	return fs.ReadFile(ls.fs, endpoint+".json")
}

func newHybridSource(dir string) (*hybridSource, error) {
	index, err := indexFromDir(dir)
	if err != nil {
		return nil, err
	}

	return &hybridSource{
		index: &inMemorySource{data: index},
		osv:   &localSource{fs: os.DirFS(dir)},
	}, nil
}

// hybridSource reads OSV entries from a local file system, but reads
// indexes from an in-memory map.
type hybridSource struct {
	index *inMemorySource
	osv   *localSource
}

func (hs *hybridSource) get(ctx context.Context, endpoint string) (_ []byte, err error) {
	derrors.Wrap(&err, "get(%s)", endpoint)

	dir, file := filepath.Split(endpoint)

	if filepath.Dir(dir) == indexDir {
		return hs.index.get(ctx, endpoint)
	}

	return hs.osv.get(ctx, file)
}

// newInMemorySource creates a new in-memory source from OSV entries.
// Adapted from x/vulndb/internal/database.go.
func newInMemorySource(entries []*osv.Entry) (*inMemorySource, error) {
	data, err := indexFromEntries(entries)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		b, err := json.Marshal(entry)
		if err != nil {
			return nil, err
		}
		data[entryEndpoint(entry.ID)] = b
	}

	return &inMemorySource{data: data}, nil
}

// inMemorySource reads databases from an in-memory map.
// Currently intended for use only in unit tests.
type inMemorySource struct {
	data map[string][]byte
}

func (db *inMemorySource) get(ctx context.Context, endpoint string) ([]byte, error) {
	b, ok := db.data[endpoint]
	if !ok {
		return nil, fmt.Errorf("no data found at endpoint %q", endpoint)
	}
	return b, nil
}
