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

// Typoscript lexer.
var Typoscript = Register(MustNewLexer(
	&Config{
		Name:      "TypoScript",
		Aliases:   []string{"typoscript"},
		Filenames: []string{"*.ts"},
		MimeTypes: []string{"text/x-typoscript"},
		DotAll:    true,
		Priority:  0.1,
	},
	typoscriptRules,
))

func typoscriptRules() Rules {
	return Rules{
		"root": {
			Include("comment"),
			Include("constant"),
			Include("html"),
			Include("label"),
			Include("whitespace"),
			Include("keywords"),
			Include("punctuation"),
			Include("operator"),
			Include("structure"),
			Include("literal"),
			Include("other"),
		},
		"keywords": {
			{`(\[)(?i)(browser|compatVersion|dayofmonth|dayofweek|dayofyear|device|ELSE|END|GLOBAL|globalString|globalVar|hostname|hour|IP|language|loginUser|loginuser|minute|month|page|PIDinRootline|PIDupinRootline|system|treeLevel|useragent|userFunc|usergroup|version)([^\]]*)(\])`, ByGroups(LiteralStringSymbol, NameConstant, Text, LiteralStringSymbol), nil},
			{`(?=[\w\-])(HTMLparser|HTMLparser_tags|addParams|cache|encapsLines|filelink|if|imageLinkWrap|imgResource|makelinks|numRows|numberFormat|parseFunc|replacement|round|select|split|stdWrap|strPad|tableStyle|tags|textStyle|typolink)(?![\w\-])`, NameFunction, nil},
			{`(?:(=?\s*<?\s+|^\s*))(cObj|field|config|content|constants|FEData|file|frameset|includeLibs|lib|page|plugin|register|resources|sitemap|sitetitle|styles|temp|tt_[^:.\s]*|types|xmlnews|INCLUDE_TYPOSCRIPT|_CSS_DEFAULT_STYLE|_DEFAULT_PI_VARS|_LOCAL_LANG)(?![\w\-])`, ByGroups(Operator, NameBuiltin), nil},
			{`(?=[\w\-])(CASE|CLEARGIF|COA|COA_INT|COBJ_ARRAY|COLUMNS|CONTENT|CTABLE|EDITPANEL|FILE|FILES|FLUIDTEMPLATE|FORM|HMENU|HRULER|HTML|IMAGE|IMGTEXT|IMG_RESOURCE|LOAD_REGISTER|MEDIA|MULTIMEDIA|OTABLE|PAGE|QTOBJECT|RECORDS|RESTORE_REGISTER|SEARCHRESULT|SVG|SWFOBJECT|TEMPLATE|TEXT|USER|USER_INT)(?![\w\-])`, NameClass, nil},
			{`(?=[\w\-])(ACTIFSUBRO|ACTIFSUB|ACTRO|ACT|CURIFSUBRO|CURIFSUB|CURRO|CUR|IFSUBRO|IFSUB|NO|SPC|USERDEF1RO|USERDEF1|USERDEF2RO|USERDEF2|USRRO|USR)`, NameClass, nil},
			{`(?=[\w\-])(GMENU_FOLDOUT|GMENU_LAYERS|GMENU|IMGMENUITEM|IMGMENU|JSMENUITEM|JSMENU|TMENUITEM|TMENU_LAYERS|TMENU)`, NameClass, nil},
			{`(?=[\w\-])(PHP_SCRIPT(_EXT|_INT)?)`, NameClass, nil},
			{`(?=[\w\-])(userFunc)(?![\w\-])`, NameFunction, nil},
		},
		"whitespace": {
			{`\s+`, Text, nil},
		},
		"html": {
			{`<\S[^\n>]*>`, Using("TypoScriptHTMLData"), nil},
			{`&[^;\n]*;`, LiteralString, nil},
			{`(_CSS_DEFAULT_STYLE)(\s*)(\()(?s)(.*(?=\n\)))`, ByGroups(NameClass, Text, LiteralStringSymbol, Using("TypoScriptCSSData")), nil},
		},
		"literal": {
			{`0x[0-9A-Fa-f]+t?`, LiteralNumberHex, nil},
			{`[0-9]+`, LiteralNumberInteger, nil},
			{`(###\w+###)`, NameConstant, nil},
		},
		"label": {
			{`(EXT|FILE|LLL):[^}\n"]*`, LiteralString, nil},
			{`(?![^\w\-])([\w\-]+(?:/[\w\-]+)+/?)(\S*\n)`, ByGroups(LiteralString, LiteralString), nil},
		},
		"punctuation": {
			{`[,.]`, Punctuation, nil},
		},
		"operator": {
			{`[<>,:=.*%+|]`, Operator, nil},
		},
		"structure": {
			{`[{}()\[\]\\]`, LiteralStringSymbol, nil},
		},
		"constant": {
			{`(\{)(\$)((?:[\w\-]+\.)*)([\w\-]+)(\})`, ByGroups(LiteralStringSymbol, Operator, NameConstant, NameConstant, LiteralStringSymbol), nil},
			{`(\{)([\w\-]+)(\s*:\s*)([\w\-]+)(\})`, ByGroups(LiteralStringSymbol, NameConstant, Operator, NameConstant, LiteralStringSymbol), nil},
			{`(#[a-fA-F0-9]{6}\b|#[a-fA-F0-9]{3}\b)`, LiteralStringChar, nil},
		},
		"comment": {
			{`(?<!(#|\'|"))(?:#(?!(?:[a-fA-F0-9]{6}|[a-fA-F0-9]{3}))[^\n#]+|//[^\n]*)`, Comment, nil},
			{`/\*(?:(?!\*/).)*\*/`, Comment, nil},
			{`(\s*#\s*\n)`, Comment, nil},
		},
		"other": {
			{`[\w"\-!/&;]+`, Text, nil},
		},
	}
}
