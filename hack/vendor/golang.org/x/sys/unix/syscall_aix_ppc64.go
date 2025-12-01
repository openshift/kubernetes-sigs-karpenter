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

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix && ppc64

package unix

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = lseek

//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) = mmap64

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: nsec}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int64(sec), Usec: int32(usec)}
}

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}

func (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint32(length)
}

func (msghdr *Msghdr) SetIovlen(length int) {
	msghdr.Iovlen = int32(length)
}

func (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}

// In order to only have Timespec structure, type of Stat_t's fields
// Atim, Mtim and Ctim is changed from StTimespec to Timespec during
// ztypes generation.
// On ppc64, Timespec.Nsec is an int64 while StTimespec.Nsec is an
// int32, so the fields' value must be modified.
func fixStatTimFields(stat *Stat_t) {
	stat.Atim.Nsec >>= 32
	stat.Mtim.Nsec >>= 32
	stat.Ctim.Nsec >>= 32
}

func Fstat(fd int, stat *Stat_t) error {
	err := fstat(fd, stat)
	if err != nil {
		return err
	}
	fixStatTimFields(stat)
	return nil
}

func Fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	err := fstatat(dirfd, path, stat, flags)
	if err != nil {
		return err
	}
	fixStatTimFields(stat)
	return nil
}

func Lstat(path string, stat *Stat_t) error {
	err := lstat(path, stat)
	if err != nil {
		return err
	}
	fixStatTimFields(stat)
	return nil
}

func Stat(path string, statptr *Stat_t) error {
	err := stat(path, statptr)
	if err != nil {
		return err
	}
	fixStatTimFields(statptr)
	return nil
}
