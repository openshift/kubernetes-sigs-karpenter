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

package ifelse

// Chain contains information about an if-else chain.
type Chain struct {
	If                   Branch     // what happens at the end of the "if" block
	HasElse              bool       // is there an "else" block?
	Else                 Branch     // what happens at the end of the "else" block
	HasInitializer       bool       // is there an "if"-initializer somewhere in the chain?
	HasPriorNonDeviating bool       // is there a prior "if" block that does NOT deviate control flow?
	AtBlockEnd           bool       // whether the chain is placed at the end of the surrounding block
	BlockEndKind         BranchKind // control flow at end of surrounding block (e.g. "return" for function body)
}
