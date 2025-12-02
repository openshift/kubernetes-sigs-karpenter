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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

func initS390Xbase() {
	// get the facilities list
	facilities := stfle()

	// mandatory
	S390X.HasZARCH = facilities.Has(zarch)
	S390X.HasSTFLE = facilities.Has(stflef)
	S390X.HasLDISP = facilities.Has(ldisp)
	S390X.HasEIMM = facilities.Has(eimm)

	// optional
	S390X.HasETF3EH = facilities.Has(etf3eh)
	S390X.HasDFP = facilities.Has(dfp)
	S390X.HasMSA = facilities.Has(msa)
	S390X.HasVX = facilities.Has(vx)
	if S390X.HasVX {
		S390X.HasVXE = facilities.Has(vxe)
	}
}
