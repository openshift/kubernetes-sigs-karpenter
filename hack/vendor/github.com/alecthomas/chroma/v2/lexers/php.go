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

package lexers

import (
	"strings"

	. "github.com/alecthomas/chroma/v2" // nolint
)

// phtml lexer is PHP in HTML.
var _ = Register(DelegatingLexer(HTML, MustNewLexer(
	&Config{
		Name:            "PHTML",
		Aliases:         []string{"phtml"},
		Filenames:       []string{"*.phtml", "*.php", "*.php[345]", "*.inc"},
		MimeTypes:       []string{"application/x-php", "application/x-httpd-php", "application/x-httpd-php3", "application/x-httpd-php4", "application/x-httpd-php5", "text/x-php"},
		DotAll:          true,
		CaseInsensitive: true,
		EnsureNL:        true,
		Priority:        2,
	},
	func() Rules {
		return Get("PHP").(*RegexLexer).MustRules().
			Rename("root", "php").
			Merge(Rules{
				"root": {
					{`<\?(php)?`, CommentPreproc, Push("php")},
					{`[^<]+`, Other, nil},
					{`<`, Other, nil},
				},
			})
	},
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "<?php") {
		return 0.5
	}
	return 0.0
})))
