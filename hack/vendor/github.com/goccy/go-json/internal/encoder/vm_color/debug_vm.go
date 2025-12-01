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

package vm_color

import (
	"fmt"

	"github.com/goccy/go-json/internal/encoder"
)

func DebugRun(ctx *encoder.RuntimeContext, b []byte, codeSet *encoder.OpcodeSet) ([]byte, error) {
	var code *encoder.Opcode
	if (ctx.Option.Flag & encoder.HTMLEscapeOption) != 0 {
		code = codeSet.EscapeKeyCode
	} else {
		code = codeSet.NoescapeKeyCode
	}

	defer func() {
		if err := recover(); err != nil {
			w := ctx.Option.DebugOut
			fmt.Fprintln(w, "=============[DEBUG]===============")
			fmt.Fprintln(w, "* [TYPE]")
			fmt.Fprintln(w, codeSet.Type)
			fmt.Fprintf(w, "\n")
			fmt.Fprintln(w, "* [ALL OPCODE]")
			fmt.Fprintln(w, code.Dump())
			fmt.Fprintf(w, "\n")
			fmt.Fprintln(w, "* [CONTEXT]")
			fmt.Fprintf(w, "%+v\n", ctx)
			fmt.Fprintln(w, "===================================")
			panic(err)
		}
	}()

	return Run(ctx, b, codeSet)
}
