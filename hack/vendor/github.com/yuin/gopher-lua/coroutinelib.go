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

package lua

func OpenCoroutine(L *LState) int {
	// TODO: Tie module name to contents of linit.go?
	mod := L.RegisterModule(CoroutineLibName, coFuncs)
	L.Push(mod)
	return 1
}

var coFuncs = map[string]LGFunction{
	"create":  coCreate,
	"yield":   coYield,
	"resume":  coResume,
	"running": coRunning,
	"status":  coStatus,
	"wrap":    coWrap,
}

func coCreate(L *LState) int {
	fn := L.CheckFunction(1)
	newthread, _ := L.NewThread()
	base := 0
	newthread.stack.Push(callFrame{
		Fn:         fn,
		Pc:         0,
		Base:       base,
		LocalBase:  base + 1,
		ReturnBase: base,
		NArgs:      0,
		NRet:       MultRet,
		Parent:     nil,
		TailCall:   0,
	})
	L.Push(newthread)
	return 1
}

func coYield(L *LState) int {
	return -1
}

func coResume(L *LState) int {
	th := L.CheckThread(1)
	if L.G.CurrentThread == th {
		msg := "can not resume a running thread"
		if th.wrapped {
			L.RaiseError(msg)
			return 0
		}
		L.Push(LFalse)
		L.Push(LString(msg))
		return 2
	}
	if th.Dead {
		msg := "can not resume a dead thread"
		if th.wrapped {
			L.RaiseError(msg)
			return 0
		}
		L.Push(LFalse)
		L.Push(LString(msg))
		return 2
	}
	th.Parent = L
	L.G.CurrentThread = th
	if !th.isStarted() {
		cf := th.stack.Last()
		th.currentFrame = cf
		th.SetTop(0)
		nargs := L.GetTop() - 1
		L.XMoveTo(th, nargs)
		cf.NArgs = nargs
		th.initCallFrame(cf)
		th.Panic = panicWithoutTraceback
	} else {
		nargs := L.GetTop() - 1
		L.XMoveTo(th, nargs)
	}
	top := L.GetTop()
	threadRun(th)
	return L.GetTop() - top
}

func coRunning(L *LState) int {
	if L.G.MainThread == L {
		L.Push(LNil)
		return 1
	}
	L.Push(L.G.CurrentThread)
	return 1
}

func coStatus(L *LState) int {
	L.Push(LString(L.Status(L.CheckThread(1))))
	return 1
}

func wrapaux(L *LState) int {
	L.Insert(L.ToThread(UpvalueIndex(1)), 1)
	return coResume(L)
}

func coWrap(L *LState) int {
	coCreate(L)
	L.CheckThread(L.GetTop()).wrapped = true
	v := L.Get(L.GetTop())
	L.Pop(1)
	L.Push(L.NewClosure(wrapaux, v))
	return 1
}

//
