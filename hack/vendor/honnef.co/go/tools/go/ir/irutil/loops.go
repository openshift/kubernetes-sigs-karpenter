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

package irutil

import "honnef.co/go/tools/go/ir"

type Loop struct{ *ir.BlockSet }

func FindLoops(fn *ir.Function) []Loop {
	if fn.Blocks == nil {
		return nil
	}
	tree := fn.DomPreorder()
	var sets []Loop
	for _, h := range tree {
		for _, n := range h.Preds {
			if !h.Dominates(n) {
				continue
			}
			// n is a back-edge to h
			// h is the loop header
			if n == h {
				set := Loop{ir.NewBlockSet(len(fn.Blocks))}
				set.Add(n)
				sets = append(sets, set)
				continue
			}
			set := Loop{ir.NewBlockSet(len(fn.Blocks))}
			set.Add(h)
			set.Add(n)
			for _, b := range allPredsBut(n, h, nil) {
				set.Add(b)
			}
			sets = append(sets, set)
		}
	}
	return sets
}

func allPredsBut(b, but *ir.BasicBlock, list []*ir.BasicBlock) []*ir.BasicBlock {
outer:
	for _, pred := range b.Preds {
		if pred == but {
			continue
		}
		for _, p := range list {
			// TODO improve big-o complexity of this function
			if pred == p {
				continue outer
			}
		}
		list = append(list, pred)
		list = allPredsBut(pred, but, list)
	}
	return list
}
