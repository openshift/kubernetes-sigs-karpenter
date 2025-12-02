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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package externalaccount

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type fileCredentialSource struct {
	File   string
	Format Format
}

func (cs fileCredentialSource) credentialSourceType() string {
	return "file"
}

func (cs fileCredentialSource) subjectToken() (string, error) {
	tokenFile, err := os.Open(cs.File)
	if err != nil {
		return "", fmt.Errorf("oauth2/google/externalaccount: failed to open credential file %q", cs.File)
	}
	defer tokenFile.Close()
	tokenBytes, err := io.ReadAll(io.LimitReader(tokenFile, 1<<20))
	if err != nil {
		return "", fmt.Errorf("oauth2/google/externalaccount: failed to read credential file: %v", err)
	}
	tokenBytes = bytes.TrimSpace(tokenBytes)
	switch cs.Format.Type {
	case "json":
		jsonData := make(map[string]any)
		err = json.Unmarshal(tokenBytes, &jsonData)
		if err != nil {
			return "", fmt.Errorf("oauth2/google/externalaccount: failed to unmarshal subject token file: %v", err)
		}
		val, ok := jsonData[cs.Format.SubjectTokenFieldName]
		if !ok {
			return "", errors.New("oauth2/google/externalaccount: provided subject_token_field_name not found in credentials")
		}
		token, ok := val.(string)
		if !ok {
			return "", errors.New("oauth2/google/externalaccount: improperly formatted subject token")
		}
		return token, nil
	case "text":
		return string(tokenBytes), nil
	case "":
		return string(tokenBytes), nil
	default:
		return "", errors.New("oauth2/google/externalaccount: invalid credential_source file format type")
	}

}
