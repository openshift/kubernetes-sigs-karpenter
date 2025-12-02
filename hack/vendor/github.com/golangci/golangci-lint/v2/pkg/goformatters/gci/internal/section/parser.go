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

package section

import (
	"errors"
	"fmt"
	"strings"

	"github.com/daixiang0/gci/pkg/section"
)

func Parse(data []string) (section.SectionList, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var list section.SectionList
	var errString string
	for _, d := range data {
		s := strings.ToLower(d)
		if len(s) == 0 {
			return nil, nil
		}

		if s == "default" {
			list = append(list, section.Default{})
		} else if s == "standard" {
			list = append(list, Standard{})
		} else if s == "newline" {
			list = append(list, section.NewLine{})
		} else if strings.HasPrefix(s, "prefix(") && len(d) > 8 {
			list = append(list, section.Custom{Prefix: d[7 : len(d)-1]})
		} else if strings.HasPrefix(s, "commentline(") && len(d) > 13 {
			list = append(list, section.Custom{Prefix: d[12 : len(d)-1]})
		} else if s == "dot" {
			list = append(list, section.Dot{})
		} else if s == "blank" {
			list = append(list, section.Blank{})
		} else if s == "alias" {
			list = append(list, section.Alias{})
		} else if s == "localmodule" {
			// pointer because we need to mutate the section at configuration time
			list = append(list, &section.LocalModule{})
		} else {
			errString += fmt.Sprintf(" %s", s)
		}
	}
	if errString != "" {
		return nil, errors.New(fmt.Sprintf("invalid params:%s", errString))
	}
	return list, nil
}
