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

package cmd

var unwrapScalarFlag = newUnwrapFlag()

var unwrapScalar = false

var writeInplace = false
var outputToJSON = false

var outputFormat = ""

var inputFormat = ""

var exitStatus = false
var indent = 2
var noDocSeparators = false
var nullInput = false
var nulSepOutput = false
var verbose = false
var version = false
var prettyPrint = false

var forceColor = false
var forceNoColor = false
var colorsEnabled = false

// can be either "" (off), "extract" or "process"
var frontMatter = ""

var splitFileExp = ""
var splitFileExpFile = ""

var completedSuccessfully = false

var forceExpression = ""

var expressionFile = ""
