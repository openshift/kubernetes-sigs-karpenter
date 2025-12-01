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

package errorlint

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type verb struct {
	format       string
	formatOffset int
	index        int
}

type printfParser struct {
	str string
	at  int
}

func (pp *printfParser) ParseAllVerbs() ([]verb, error) {
	verbs := []verb{}
	for {
		verb, err := pp.parseVerb()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		verbs = append(verbs, *verb)
	}
	return verbs, nil
}

func (pp *printfParser) parseVerb() (*verb, error) {
	if err := pp.skipToPercent(); err != nil {
		return nil, err
	}
	if pp.next() != '%' {
		return nil, fmt.Errorf("expected '%%'")
	}

	index := -1
	for {
		switch pp.peek() {
		case '%':
			pp.next()
			return pp.parseVerb()
		case '+', '#':
			pp.next()
			continue
		case '[':
			var err error
			index, err = pp.parseIndex()
			if err != nil {
				return nil, err
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			pp.parsePrecision()
		case 0:
			return nil, io.EOF
		}
		break
	}

	format := pp.next()

	return &verb{format: string(format), formatOffset: pp.at - 1, index: index}, nil
}

func (pp *printfParser) parseIndex() (int, error) {
	if pp.next() != '[' {
		return -1, fmt.Errorf("expected '['")
	}
	end := strings.Index(pp.str, "]")
	if end == -1 {
		return -1, fmt.Errorf("unterminated indexed verb")
	}
	index, err := strconv.Atoi(pp.str[:end])
	if err != nil {
		return -1, err
	}
	pp.str = pp.str[end+1:]
	pp.at += end + 1
	return index, nil
}

func (pp *printfParser) parsePrecision() {
	for {
		if r := pp.peek(); (r < '0' || '9' < r) && r != '.' {
			break
		}
		pp.next()
	}
}

func (pp *printfParser) skipToPercent() error {
	i := strings.Index(pp.str, "%")
	if i == -1 {
		return io.EOF
	}
	pp.str = pp.str[i:]
	pp.at += i
	return nil
}

func (pp *printfParser) peek() rune {
	if len(pp.str) == 0 {
		return 0
	}
	return rune(pp.str[0])
}

func (pp *printfParser) next() rune {
	if len(pp.str) == 0 {
		return 0
	}
	r := rune(pp.str[0])
	pp.str = pp.str[1:]
	pp.at++
	return r
}
