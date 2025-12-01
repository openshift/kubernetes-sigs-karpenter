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

package keys

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

type PKIXPublicKey struct {
	crypto.PublicKey
}

func (p *PKIXPublicKey) MarshalJSON() ([]byte, error) {
	bytes, err := x509.MarshalPKIXPublicKey(p.PublicKey)
	if err != nil {
		return nil, err
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	})
	return json.Marshal(string(pemBytes))
}

func (p *PKIXPublicKey) UnmarshalJSON(b []byte) error {
	var pemValue string
	// Prepare decoder limited to 512Kb
	dec := json.NewDecoder(io.LimitReader(bytes.NewReader(b), MaxJSONKeySize))

	// Unmarshal key value
	if err := dec.Decode(&pemValue); err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return fmt.Errorf("tuf: the public key is truncated or too large: %w", err)
		}
		return err
	}

	block, _ := pem.Decode([]byte(pemValue))
	if block == nil {
		return errors.New("invalid PEM value")
	}
	if block.Type != "PUBLIC KEY" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	p.PublicKey = pub
	return nil
}
