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

//go:build windows

package internal

import (
	"errors"

	"golang.org/x/sys/windows"
)

// Just a dummy.
var (
	ErrSyscallEACCES = errors.New("dummy")
	ErrUnixEACCES    = errors.New("dummy")
)

func SetRlimit()                                    {}
func Maxfiles() uint64                              { return 1<<64 - 1 }
func Mkfifo(path string, mode uint32) error         { return errors.New("no FIFOs on Windows") }
func Mknod(path string, mode uint32, dev int) error { return errors.New("no device nodes on Windows") }

func HasPrivilegesForSymlink() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}
	defer windows.FreeSid(sid)
	token := windows.Token(0)
	member, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return member || token.IsElevated()
}
