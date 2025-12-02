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

//go:build !windows && !darwin && !freebsd && !plan9

package internal

import (
	"syscall"

	"golang.org/x/sys/unix"
)

var (
	ErrSyscallEACCES = syscall.EACCES
	ErrUnixEACCES    = unix.EACCES
)

var maxfiles uint64

func SetRlimit() {
	// Go 1.19 will do this automatically: https://go-review.googlesource.com/c/go/+/393354/
	var l syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &l)
	if err == nil && l.Cur != l.Max {
		l.Cur = l.Max
		syscall.Setrlimit(syscall.RLIMIT_NOFILE, &l)
	}
	maxfiles = uint64(l.Cur)
}

func Maxfiles() uint64                              { return maxfiles }
func Mkfifo(path string, mode uint32) error         { return unix.Mkfifo(path, mode) }
func Mknod(path string, mode uint32, dev int) error { return unix.Mknod(path, mode, dev) }
