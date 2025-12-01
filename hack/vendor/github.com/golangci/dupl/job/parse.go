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

package job

import (
	"log"

	"github.com/golangci/dupl/syntax"
	"github.com/golangci/dupl/syntax/golang"
)

func Parse(fchan chan string) chan []*syntax.Node {

	// parse AST
	achan := make(chan *syntax.Node)
	go func() {
		for file := range fchan {
			ast, err := golang.Parse(file)
			if err != nil {
				log.Println(err)
				continue
			}
			achan <- ast
		}
		close(achan)
	}()

	// serialize
	schan := make(chan []*syntax.Node)
	go func() {
		for ast := range achan {
			seq := syntax.Serialize(ast)
			schan <- seq
		}
		close(schan)
	}()
	return schan
}
