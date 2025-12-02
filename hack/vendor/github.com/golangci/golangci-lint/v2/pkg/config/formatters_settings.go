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

package config

var defaultFormatterSettings = FormatterSettings{
	GoFmt: GoFmtSettings{
		Simplify: true,
	},
	Gci: GciSettings{
		Sections: []string{"standard", "default"},
	},
	GoLines: GoLinesSettings{
		MaxLen:         100,
		TabLen:         4,
		ReformatTags:   true,
		ChainSplitDots: true,
	},
}

type FormatterSettings struct {
	Gci       GciSettings       `mapstructure:"gci"`
	GoFmt     GoFmtSettings     `mapstructure:"gofmt"`
	GoFumpt   GoFumptSettings   `mapstructure:"gofumpt"`
	GoImports GoImportsSettings `mapstructure:"goimports"`
	GoLines   GoLinesSettings   `mapstructure:"golines"`
}

type GciSettings struct {
	Sections         []string `mapstructure:"sections"`
	NoInlineComments bool     `mapstructure:"no-inline-comments"`
	NoPrefixComments bool     `mapstructure:"no-prefix-comments"`
	CustomOrder      bool     `mapstructure:"custom-order"`
	NoLexOrder       bool     `mapstructure:"no-lex-order"`
}

type GoFmtSettings struct {
	Simplify     bool               `mapstructure:"simplify"`
	RewriteRules []GoFmtRewriteRule `mapstructure:"rewrite-rules"`
}

type GoFmtRewriteRule struct {
	Pattern     string `mapstructure:"pattern"`
	Replacement string `mapstructure:"replacement"`
}

type GoFumptSettings struct {
	ModulePath string `mapstructure:"module-path"`
	ExtraRules bool   `mapstructure:"extra-rules"`

	LangVersion string `mapstructure:"-"`
}

type GoImportsSettings struct {
	LocalPrefixes []string `mapstructure:"local-prefixes"`
}

type GoLinesSettings struct {
	MaxLen          int  `mapstructure:"max-len"`
	TabLen          int  `mapstructure:"tab-len"`
	ShortenComments bool `mapstructure:"shorten-comments"`
	ReformatTags    bool `mapstructure:"reformat-tags"`
	ChainSplitDots  bool `mapstructure:"chain-split-dots"`
}
