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

package flect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	loadCustomData("inflections.json", "INFLECT_PATH", "could not read inflection file", LoadInflections)
	loadCustomData("acronyms.json", "ACRONYMS_PATH", "could not read acronyms file", LoadAcronyms)
}

//CustomDataParser are functions that parse data like acronyms or
//plurals in the shape of a io.Reader it receives.
type CustomDataParser func(io.Reader) error

func loadCustomData(defaultFile, env, readErrorMessage string, parser CustomDataParser) {
	pwd, _ := os.Getwd()
	path, found := os.LookupEnv(env)
	if !found {
		path = filepath.Join(pwd, defaultFile)
	}

	if _, err := os.Stat(path); err != nil {
		return
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s %s (%s)\n", readErrorMessage, path, err)
		return
	}

	if err = parser(bytes.NewReader(b)); err != nil {
		fmt.Println(err)
	}
}

//LoadAcronyms loads rules from io.Reader param
func LoadAcronyms(r io.Reader) error {
	m := []string{}
	err := json.NewDecoder(r).Decode(&m)

	if err != nil {
		return fmt.Errorf("could not decode acronyms JSON from reader: %s", err)
	}

	acronymsMoot.Lock()
	defer acronymsMoot.Unlock()

	for _, acronym := range m {
		baseAcronyms[acronym] = true
	}

	return nil
}

//LoadInflections loads rules from io.Reader param
func LoadInflections(r io.Reader) error {
	m := map[string]string{}

	err := json.NewDecoder(r).Decode(&m)
	if err != nil {
		return fmt.Errorf("could not decode inflection JSON from reader: %s", err)
	}

	pluralMoot.Lock()
	defer pluralMoot.Unlock()
	singularMoot.Lock()
	defer singularMoot.Unlock()

	for s, p := range m {
		if strings.Contains(s, " ") || strings.Contains(p, " ") {
			// flect works with parts, so multi-words should not be allowed
			return fmt.Errorf("inflection elements should be a single word")
		}
		singleToPlural[s] = p
		pluralToSingle[p] = s
	}

	return nil
}
