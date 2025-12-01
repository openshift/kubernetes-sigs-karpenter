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

package colorable

import (
	"bytes"
	"io"
)

// NonColorable holds writer but removes escape sequence.
type NonColorable struct {
	out io.Writer
}

// NewNonColorable returns new instance of Writer which removes escape sequence from Writer.
func NewNonColorable(w io.Writer) io.Writer {
	return &NonColorable{out: w}
}

// Write writes data on console
func (w *NonColorable) Write(data []byte) (n int, err error) {
	er := bytes.NewReader(data)
	var plaintext bytes.Buffer
loop:
	for {
		c1, err := er.ReadByte()
		if err != nil {
			plaintext.WriteTo(w.out)
			break loop
		}
		if c1 != 0x1b {
			plaintext.WriteByte(c1)
			continue
		}
		_, err = plaintext.WriteTo(w.out)
		if err != nil {
			break loop
		}
		c2, err := er.ReadByte()
		if err != nil {
			break loop
		}
		if c2 != 0x5b {
			continue
		}

		for {
			c, err := er.ReadByte()
			if err != nil {
				break loop
			}
			if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '@' {
				break
			}
		}
	}

	return len(data), nil
}
