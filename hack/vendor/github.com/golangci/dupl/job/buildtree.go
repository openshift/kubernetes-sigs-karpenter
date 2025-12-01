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
	"github.com/golangci/dupl/suffixtree"
	"github.com/golangci/dupl/syntax"
)

func BuildTree(schan chan []*syntax.Node) (t *suffixtree.STree, d *[]*syntax.Node, done chan bool) {
	t = suffixtree.New()
	data := make([]*syntax.Node, 0, 100)
	done = make(chan bool)
	go func() {
		for seq := range schan {
			data = append(data, seq...)
			for _, node := range seq {
				t.Update(node)
			}
		}
		done <- true
	}()
	return t, &data, done
}
