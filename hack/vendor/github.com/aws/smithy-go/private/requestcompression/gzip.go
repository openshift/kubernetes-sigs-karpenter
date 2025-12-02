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

package requestcompression

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func gzipCompress(input io.Reader) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.DefaultCompression)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip writer, %v", err)
	}

	inBytes, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("failed read payload to compress, %v", err)
	}

	if _, err = w.Write(inBytes); err != nil {
		return nil, fmt.Errorf("failed to write payload to be compressed, %v", err)
	}
	if err = w.Close(); err != nil {
		return nil, fmt.Errorf("failed to flush payload being compressed, %v", err)
	}

	return b.Bytes(), nil
}
