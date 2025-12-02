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

package yqlib

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
)

type base64Padder struct {
	count int
	io.Reader
}

func (c *base64Padder) pad(buf []byte) (int, error) {
	pad := strings.Repeat("=", (4 - c.count%4))
	n, err := strings.NewReader(pad).Read(buf)
	c.count += n
	return n, err
}

func (c *base64Padder) Read(buf []byte) (int, error) {
	n, err := c.Reader.Read(buf)
	c.count += n

	if err == io.EOF && c.count%4 != 0 {
		return c.pad(buf)
	}
	return n, err
}

type base64Decoder struct {
	reader       io.Reader
	finished     bool
	readAnything bool
	encoding     base64.Encoding
}

func NewBase64Decoder() Decoder {
	return &base64Decoder{finished: false, encoding: *base64.StdEncoding}
}

func (dec *base64Decoder) Init(reader io.Reader) error {
	dec.reader = &base64Padder{Reader: reader}
	dec.readAnything = false
	dec.finished = false
	return nil
}

func (dec *base64Decoder) Decode() (*CandidateNode, error) {
	if dec.finished {
		return nil, io.EOF
	}
	base64Reader := base64.NewDecoder(&dec.encoding, dec.reader)
	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(base64Reader); err != nil {
		return nil, err
	}
	if buf.Len() == 0 {
		dec.finished = true

		// if we've read _only_ an empty string, lets return that
		// otherwise if we've already read some bytes, and now we get
		// an empty string, then we are done.
		if dec.readAnything {
			return nil, io.EOF
		}
	}
	dec.readAnything = true
	return createStringScalarNode(buf.String()), nil
}
