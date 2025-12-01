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

package kitty

import (
	"compress/zlib"
	"fmt"
	"image"
	"image/png"
	"io"
)

// Encoder is an encoder for the Kitty graphics protocol. It supports encoding
// images in the 24-bit [RGB], 32-bit [RGBA], and [PNG] formats, and
// compressing the data using zlib.
// The default format is 32-bit [RGBA].
type Encoder struct {
	// Uses zlib compression.
	Compress bool

	// Can be one of [RGBA], [RGB], or [PNG].
	Format int
}

// Encode encodes the image data in the specified format and writes it to w.
func (e *Encoder) Encode(w io.Writer, m image.Image) error {
	if m == nil {
		return nil
	}

	if e.Compress {
		zw := zlib.NewWriter(w)
		defer zw.Close() //nolint:errcheck
		w = zw
	}

	if e.Format == 0 {
		e.Format = RGBA
	}

	switch e.Format {
	case RGBA, RGB:
		bounds := m.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				switch e.Format {
				case RGBA:
					w.Write([]byte{byte(r >> 8), byte(g >> 8), byte(b >> 8), byte(a >> 8)}) //nolint:errcheck
				case RGB:
					w.Write([]byte{byte(r >> 8), byte(g >> 8), byte(b >> 8)}) //nolint:errcheck
				}
			}
		}

	case PNG:
		if err := png.Encode(w, m); err != nil {
			return fmt.Errorf("failed to encode PNG: %w", err)
		}

	default:
		return fmt.Errorf("unsupported format: %d", e.Format)
	}

	return nil
}
