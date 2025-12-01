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

package ansi

import (
	"sync"

	"github.com/charmbracelet/x/ansi/parser"
)

var parserPool = sync.Pool{
	New: func() any {
		p := NewParser()
		p.SetParamsSize(parser.MaxParamsSize)
		p.SetDataSize(1024 * 1024 * 4) // 4MB of data buffer
		return p
	},
}

// GetParser returns a parser from a sync pool.
func GetParser() *Parser {
	return parserPool.Get().(*Parser)
}

// PutParser returns a parser to a sync pool. The parser is reset
// automatically.
func PutParser(p *Parser) {
	p.Reset()
	p.dataLen = 0
	parserPool.Put(p)
}
