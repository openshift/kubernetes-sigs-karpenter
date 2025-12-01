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

package quasigo

import (
	"fmt"
	"strings"
)

// TODO(quasilyte): generate extra opcode info so we can simplify disasm function?

func disasm(env *Env, fn *Func) string {
	var out strings.Builder

	dbg, ok := env.debug.funcs[fn]
	if !ok {
		return "<unknown>\n"
	}

	code := fn.code
	labels := map[int]string{}
	walkBytecode(code, func(pc int, op opcode) {
		switch op {
		case opJumpTrue, opJumpFalse, opJump:
			offset := decode16(code, pc+1)
			targetPC := pc + offset
			if _, ok := labels[targetPC]; !ok {
				labels[targetPC] = fmt.Sprintf("L%d", len(labels))
			}
		}
	})

	walkBytecode(code, func(pc int, op opcode) {
		if l := labels[pc]; l != "" {
			fmt.Fprintf(&out, "%s:\n", l)
		}
		var arg interface{}
		var comment string
		switch op {
		case opCallNative:
			id := decode16(code, pc+1)
			arg = id
			comment = env.nativeFuncs[id].name
		case opCall, opIntCall, opVoidCall:
			id := decode16(code, pc+1)
			arg = id
			comment = env.userFuncs[id].name
		case opPushParam:
			index := int(code[pc+1])
			arg = index
			comment = dbg.paramNames[index]
		case opPushIntParam:
			index := int(code[pc+1])
			arg = index
			comment = dbg.intParamNames[index]
		case opSetLocal, opSetIntLocal, opPushLocal, opPushIntLocal, opIncLocal, opDecLocal:
			index := int(code[pc+1])
			arg = index
			comment = dbg.localNames[index]
		case opSetVariadicLen:
			arg = int(code[pc+1])
		case opPushConst:
			arg = int(code[pc+1])
			comment = fmt.Sprintf("value=%#v", fn.constants[code[pc+1]])
		case opPushIntConst:
			arg = int(code[pc+1])
			comment = fmt.Sprintf("value=%#v", fn.intConstants[code[pc+1]])
		case opJumpTrue, opJumpFalse, opJump:
			offset := decode16(code, pc+1)
			targetPC := pc + offset
			arg = offset
			comment = labels[targetPC]
		}

		if comment != "" {
			comment = " # " + comment
		}
		if arg == nil {
			fmt.Fprintf(&out, "  %s%s\n", op, comment)
		} else {
			fmt.Fprintf(&out, "  %s %#v%s\n", op, arg, comment)
		}
	})

	return out.String()
}
