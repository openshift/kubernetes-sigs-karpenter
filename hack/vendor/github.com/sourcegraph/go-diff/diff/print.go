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

package diff

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"time"
)

// PrintMultiFileDiff prints a multi-file diff in unified diff format.
func PrintMultiFileDiff(ds []*FileDiff) ([]byte, error) {
	var buf bytes.Buffer
	for _, d := range ds {
		diff, err := PrintFileDiff(d)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(diff); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// PrintFileDiff prints a FileDiff in unified diff format.
//
// TODO(sqs): handle escaping whitespace/etc. chars in filenames
func PrintFileDiff(d *FileDiff) ([]byte, error) {
	var buf bytes.Buffer

	for _, xheader := range d.Extended {
		if _, err := fmt.Fprintln(&buf, xheader); err != nil {
			return nil, err
		}
	}

	// FileDiff is added/deleted file
	// No further hunks printing needed
	if d.NewName == "" {
		_, err := fmt.Fprintf(&buf, onlyInMessage, filepath.Dir(d.OrigName), filepath.Base(d.OrigName))
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	if d.Hunks == nil {
		return buf.Bytes(), nil
	}

	if err := printFileHeader(&buf, "--- ", d.OrigName, d.OrigTime); err != nil {
		return nil, err
	}
	if err := printFileHeader(&buf, "+++ ", d.NewName, d.NewTime); err != nil {
		return nil, err
	}

	ph, err := PrintHunks(d.Hunks)
	if err != nil {
		return nil, err
	}

	if _, err := buf.Write(ph); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func printFileHeader(w io.Writer, prefix string, filename string, timestamp *time.Time) error {
	if _, err := fmt.Fprint(w, prefix, filename); err != nil {
		return err
	}
	if timestamp != nil {
		if _, err := fmt.Fprint(w, "\t", timestamp.Format(diffTimeFormatLayout)); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	return nil
}

// PrintHunks prints diff hunks in unified diff format.
func PrintHunks(hunks []*Hunk) ([]byte, error) {
	var buf bytes.Buffer
	for _, hunk := range hunks {
		_, err := fmt.Fprintf(&buf,
			"@@ -%d,%d +%d,%d @@", hunk.OrigStartLine, hunk.OrigLines, hunk.NewStartLine, hunk.NewLines,
		)
		if err != nil {
			return nil, err
		}
		if hunk.Section != "" {
			_, err := fmt.Fprint(&buf, " ", hunk.Section)
			if err != nil {
				return nil, err
			}
		}
		if _, err := fmt.Fprintln(&buf); err != nil {
			return nil, err
		}

		if hunk.OrigNoNewlineAt == 0 {
			if _, err := buf.Write(hunk.Body); err != nil {
				return nil, err
			}
		} else {
			if _, err := buf.Write(hunk.Body[:hunk.OrigNoNewlineAt]); err != nil {
				return nil, err
			}
			if err := printNoNewlineMessage(&buf); err != nil {
				return nil, err
			}
			if _, err := buf.Write(hunk.Body[hunk.OrigNoNewlineAt:]); err != nil {
				return nil, err
			}
		}

		if !bytes.HasSuffix(hunk.Body, []byte{'\n'}) {
			if _, err := fmt.Fprintln(&buf); err != nil {
				return nil, err
			}
			if err := printNoNewlineMessage(&buf); err != nil {
				return nil, err
			}
		}
	}
	return buf.Bytes(), nil
}

func printNoNewlineMessage(w io.Writer) error {
	if _, err := w.Write([]byte(noNewlineMessage)); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	return nil
}
