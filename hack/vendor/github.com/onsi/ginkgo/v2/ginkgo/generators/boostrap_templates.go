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

package generators

var bootstrapText = `package {{.Package}}

import (
	"testing"

	{{.GinkgoImport}}
	{{.GomegaImport}}
)

func Test{{.FormattedName}}(t *testing.T) {
	{{.GomegaPackage}}RegisterFailHandler({{.GinkgoPackage}}Fail)
	{{.GinkgoPackage}}RunSpecs(t, "{{.FormattedName}} Suite")
}
`

var agoutiBootstrapText = `package {{.Package}}

import (
	"testing"

	{{.GinkgoImport}}
	{{.GomegaImport}}
	"github.com/sclevine/agouti"
)

func Test{{.FormattedName}}(t *testing.T) {
	{{.GomegaPackage}}RegisterFailHandler({{.GinkgoPackage}}Fail)
	{{.GinkgoPackage}}RunSpecs(t, "{{.FormattedName}} Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = {{.GinkgoPackage}}BeforeSuite(func() {
	// Choose a WebDriver:

	agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	// agoutiDriver = agouti.ChromeDriver()

	{{.GomegaPackage}}Expect(agoutiDriver.Start()).To({{.GomegaPackage}}Succeed())
})

var _ = {{.GinkgoPackage}}AfterSuite(func() {
	{{.GomegaPackage}}Expect(agoutiDriver.Stop()).To({{.GomegaPackage}}Succeed())
})
`
