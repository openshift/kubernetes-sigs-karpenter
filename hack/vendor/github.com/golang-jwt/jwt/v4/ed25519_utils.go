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

package jwt

import (
	"crypto"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrNotEdPrivateKey = errors.New("key is not a valid Ed25519 private key")
	ErrNotEdPublicKey  = errors.New("key is not a valid Ed25519 public key")
)

// ParseEdPrivateKeyFromPEM parses a PEM-encoded Edwards curve private key
func ParseEdPrivateKeyFromPEM(key []byte) (crypto.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey ed25519.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(ed25519.PrivateKey); !ok {
		return nil, ErrNotEdPrivateKey
	}

	return pkey, nil
}

// ParseEdPublicKeyFromPEM parses a PEM-encoded Edwards curve public key
func ParseEdPublicKeyFromPEM(key []byte) (crypto.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey ed25519.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(ed25519.PublicKey); !ok {
		return nil, ErrNotEdPublicKey
	}

	return pkey, nil
}
