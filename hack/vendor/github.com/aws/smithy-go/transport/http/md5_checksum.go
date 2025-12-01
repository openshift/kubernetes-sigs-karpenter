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

package http

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

// computeMD5Checksum computes base64 md5 checksum of an io.Reader's contents.
// Returns the byte slice of md5 checksum and an error.
func computeMD5Checksum(r io.Reader) ([]byte, error) {
	h := md5.New()
	// copy errors may be assumed to be from the body.
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	// encode the md5 checksum in base64.
	sum := h.Sum(nil)
	sum64 := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
	base64.StdEncoding.Encode(sum64, sum)
	return sum64, nil
}
