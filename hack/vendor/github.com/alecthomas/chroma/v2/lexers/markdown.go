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

// Markdown lexer.
var Markdown = Register(MustNewLexer(
	&Config{
		Name:      "markdown",
		Aliases:   []string{"md", "mkd"},
		Filenames: []string{"*.md", "*.mkd", "*.markdown"},
		MimeTypes: []string{"text/x-markdown"},
	},
	markdownRules,
))

func markdownRules() Rules {
	return Rules{
		"root": {
			{`^(#[^#].+\n)`, ByGroups(GenericHeading), nil},
			{`^(#{2,6}.+\n)`, ByGroups(GenericSubheading), nil},
			{`^(\s*)([*-] )(\[[ xX]\])( .+\n)`, ByGroups(Text, Keyword, Keyword, UsingSelf("inline")), nil},
			{`^(\s*)([*-])(\s)(.+\n)`, ByGroups(Text, Keyword, Text, UsingSelf("inline")), nil},
			{`^(\s*)([0-9]+\.)( .+\n)`, ByGroups(Text, Keyword, UsingSelf("inline")), nil},
			{`^(\s*>\s)(.+\n)`, ByGroups(Keyword, GenericEmph), nil},
			{"^(```\\n)([\\w\\W]*?)(^```$)", ByGroups(String, Text, String), nil},
			{
				"^(```)(\\w+)(\\n)([\\w\\W]*?)(^```$)",
				UsingByGroup(2, 4, String, String, String, Text, String),
				nil,
			},
			Include("inline"),
		},
		"inline": {
			{`\\.`, Text, nil},
			{`(\s)(\*|_)((?:(?!\2).)*)(\2)((?=\W|\n))`, ByGroups(Text, GenericEmph, GenericEmph, GenericEmph, Text), nil},
			{`(\s)((\*\*|__).*?)\3((?=\W|\n))`, ByGroups(Text, GenericStrong, GenericStrong, Text), nil},
			{`(\s)(~~[^~]+~~)((?=\W|\n))`, ByGroups(Text, GenericDeleted, Text), nil},
			{"`[^`]+`", LiteralStringBacktick, nil},
			{`[@#][\w/:]+`, NameEntity, nil},
			{`(!?\[)([^]]+)(\])(\()([^)]+)(\))`, ByGroups(Text, NameTag, Text, Text, NameAttribute, Text), nil},
			{`.|\n`, Text, nil},
		},
	}
}
