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

package config

// GinkgoConfigType has been deprecated and its equivalent now lives in
// the types package.  You can no longer access Ginkgo configuration from the config
// package.  Instead use the DSL's GinkgoConfiguration() function to get copies of the
// current configuration
//
// GinkgoConfigType is still here so custom V1 reporters do not result in a compilation error
// It will be removed in a future minor release of Ginkgo
type GinkgoConfigType = DeprecatedGinkgoConfigType
type DeprecatedGinkgoConfigType struct {
	RandomSeed         int64
	RandomizeAllSpecs  bool
	RegexScansFilePath bool
	FocusStrings       []string
	SkipStrings        []string
	SkipMeasurements   bool
	FailOnPending      bool
	FailFast           bool
	FlakeAttempts      int
	EmitSpecProgress   bool
	DryRun             bool
	DebugParallel      bool

	ParallelNode  int
	ParallelTotal int
	SyncHost      string
	StreamHost    string
}

// DefaultReporterConfigType has been deprecated and its equivalent now lives in
// the types package.  You can no longer access Ginkgo configuration from the config
// package.  Instead use the DSL's GinkgoConfiguration() function to get copies of the
// current configuration
//
// DefaultReporterConfigType is still here so custom V1 reporters do not result in a compilation error
// It will be removed in a future minor release of Ginkgo
type DefaultReporterConfigType = DeprecatedDefaultReporterConfigType
type DeprecatedDefaultReporterConfigType struct {
	NoColor           bool
	SlowSpecThreshold float64
	NoisyPendings     bool
	NoisySkippings    bool
	Succinct          bool
	Verbose           bool
	FullTrace         bool
	ReportPassed      bool
	ReportFile        string
}

// Sadly there is no way to gracefully deprecate access to these global config variables.
// Users who need access to Ginkgo's configuration should use the DSL's GinkgoConfiguration() method
// These new unwieldy type names exist to give users a hint when they try to compile and the compilation fails
type GinkgoConfigIsNoLongerAccessibleFromTheConfigPackageUseTheDSLsGinkgoConfigurationFunctionInstead struct{}

// Sadly there is no way to gracefully deprecate access to these global config variables.
// Users who need access to Ginkgo's configuration should use the DSL's GinkgoConfiguration() method
// These new unwieldy type names exist to give users a hint when they try to compile and the compilation fails
var GinkgoConfig = GinkgoConfigIsNoLongerAccessibleFromTheConfigPackageUseTheDSLsGinkgoConfigurationFunctionInstead{}

// Sadly there is no way to gracefully deprecate access to these global config variables.
// Users who need access to Ginkgo's configuration should use the DSL's GinkgoConfiguration() method
// These new unwieldy type names exist to give users a hint when they try to compile and the compilation fails
type DefaultReporterConfigIsNoLongerAccessibleFromTheConfigPackageUseTheDSLsGinkgoConfigurationFunctionInstead struct{}

// Sadly there is no way to gracefully deprecate access to these global config variables.
// Users who need access to Ginkgo's configuration should use the DSL's GinkgoConfiguration() method
// These new unwieldy type names exist to give users a hint when they try to compile and the compilation fails
var DefaultReporterConfig = DefaultReporterConfigIsNoLongerAccessibleFromTheConfigPackageUseTheDSLsGinkgoConfigurationFunctionInstead{}
