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

package gogrep

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func compileImportPattern(config CompileConfig) (*Pattern, PatternInfo, error) {
	// TODO: figure out how to compile it as a part of a normal pattern compilation?
	// This is an adhoc solution to a problem.

	readIdent := func(s string) (varname, rest string) {
		first := true
		var offset int
		for _, ch := range s {
			ok := unicode.IsLetter(ch) ||
				ch == '_' ||
				(!first && unicode.IsDigit(ch))
			if !ok {
				break
			}
			offset += utf8.RuneLen(ch)
			first = false
		}
		return s[:offset], s[offset:]
	}

	info := newPatternInfo()
	src := config.Src
	src = src[len("import $"):]
	if src == "" {
		return nil, info, errors.New("expected ident after $, found EOF")
	}
	varname, rest := readIdent(src)
	if strings.TrimSpace(rest) != "" {
		return nil, info, fmt.Errorf("unexpected %s", rest)
	}
	var p program
	if varname != "_" {
		info.Vars[src] = struct{}{}
		p.strings = []string{varname}
		p.insts = []instruction{
			{op: opImportDecl},
			{op: opNamedNodeSeq, valueIndex: 0},
			{op: opEnd},
		}
	} else {
		p.insts = []instruction{
			{op: opAnyImportDecl},
		}
	}
	m := matcher{prog: &p, insts: p.insts}
	return &Pattern{m: &m}, info, nil
}
