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
	. "github.com/alecthomas/chroma/v2" // nolint
)

// Gemtext lexer.
var Gemtext = Register(MustNewLexer(
	&Config{
		Name:      "Gemtext",
		Aliases:   []string{"gemtext", "gmi", "gmni", "gemini"},
		Filenames: []string{"*.gmi", "*.gmni", "*.gemini"},
		MimeTypes: []string{"text/gemini"},
	},
	gemtextRules,
))

func gemtextRules() Rules {
	return Rules{
		"root": {
			{`^(#[^#].+\r?\n)`, ByGroups(GenericHeading), nil},
			{`^(#{2,3}.+\r?\n)`, ByGroups(GenericSubheading), nil},
			{`^(\* )(.+\r?\n)`, ByGroups(Keyword, Text), nil},
			{`^(>)(.+\r?\n)`, ByGroups(Keyword, GenericEmph), nil},
			{"^(```\\r?\\n)([\\w\\W]*?)(^```)(.+\\r?\\n)?", ByGroups(String, Text, String, Comment), nil},
			{
				"^(```)(\\w+)(\\r?\\n)([\\w\\W]*?)(^```)(.+\\r?\\n)?",
				UsingByGroup(2, 4, String, String, String, Text, String, Comment),
				nil,
			},
			{"^(```)(.+\\r?\\n)([\\w\\W]*?)(^```)(.+\\r?\\n)?", ByGroups(String, String, Text, String, Comment), nil},
			{`^(=>)(\s*)([^\s]+)(\s*)$`, ByGroups(Keyword, Text, NameAttribute, Text), nil},
			{`^(=>)(\s*)([^\s]+)(\s+)(.+)$`, ByGroups(Keyword, Text, NameAttribute, Text, NameTag), nil},
			{`(.|(?:\r?\n))`, ByGroups(Text), nil},
		},
	}
}
