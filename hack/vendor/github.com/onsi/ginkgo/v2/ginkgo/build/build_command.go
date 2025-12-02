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

package build

import (
	"fmt"
	"os"
	"path"

	"github.com/onsi/ginkgo/v2/ginkgo/command"
	"github.com/onsi/ginkgo/v2/ginkgo/internal"
	"github.com/onsi/ginkgo/v2/types"
)

func BuildBuildCommand() command.Command {
	var cliConfig = types.NewDefaultCLIConfig()
	var goFlagsConfig = types.NewDefaultGoFlagsConfig()

	flags, err := types.BuildBuildCommandFlagSet(&cliConfig, &goFlagsConfig)
	if err != nil {
		panic(err)
	}

	return command.Command{
		Name:     "build",
		Flags:    flags,
		Usage:    "ginkgo build <FLAGS> <PACKAGES>",
		ShortDoc: "Build the passed in <PACKAGES> (or the package in the current directory if left blank).",
		DocLink:  "precompiling-suites",
		Command: func(args []string, _ []string) {
			var errors []error
			cliConfig, goFlagsConfig, errors = types.VetAndInitializeCLIAndGoConfig(cliConfig, goFlagsConfig)
			command.AbortIfErrors("Ginkgo detected configuration issues:", errors)
			buildSpecs(args, cliConfig, goFlagsConfig)
		},
	}
}

func buildSpecs(args []string, cliConfig types.CLIConfig, goFlagsConfig types.GoFlagsConfig) {
	suites := internal.FindSuites(args, cliConfig, false).WithoutState(internal.TestSuiteStateSkippedByFilter)
	if len(suites) == 0 {
		command.AbortWith("Found no test suites")
	}

	internal.VerifyCLIAndFrameworkVersion(suites)

	opc := internal.NewOrderedParallelCompiler(cliConfig.ComputedNumCompilers())
	opc.StartCompiling(suites, goFlagsConfig, true)

	for {
		suiteIdx, suite := opc.Next()
		if suiteIdx >= len(suites) {
			break
		}
		suites[suiteIdx] = suite
		if suite.State.Is(internal.TestSuiteStateFailedToCompile) {
			fmt.Println(suite.CompilationError.Error())
		} else {
			var testBinPath string
			if len(goFlagsConfig.O) != 0 {
				stat, err := os.Stat(goFlagsConfig.O)
				if err != nil {
					panic(err)
				}
				if stat.IsDir() {
					testBinPath = goFlagsConfig.O + "/" + suite.PackageName + ".test"
				} else {
					testBinPath = goFlagsConfig.O
				}
			}
			if len(testBinPath) == 0 {
				testBinPath = path.Join(suite.Path, suite.PackageName+".test")
			}
			fmt.Printf("Compiled %s\n", testBinPath)
		}
	}

	if suites.CountWithState(internal.TestSuiteStateFailedToCompile) > 0 {
		command.AbortWith("Failed to compile all tests")
	}
}
