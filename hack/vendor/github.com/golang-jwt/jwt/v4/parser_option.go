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

package jwt

// ParserOption is used to implement functional-style options that modify the behavior of the parser. To add
// new options, just create a function (ideally beginning with With or Without) that returns an anonymous function that
// takes a *Parser type as input and manipulates its configuration accordingly.
type ParserOption func(*Parser)

// WithValidMethods is an option to supply algorithm methods that the parser will check. Only those methods will be considered valid.
// It is heavily encouraged to use this option in order to prevent attacks such as https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/.
func WithValidMethods(methods []string) ParserOption {
	return func(p *Parser) {
		p.ValidMethods = methods
	}
}

// WithJSONNumber is an option to configure the underlying JSON parser with UseNumber
func WithJSONNumber() ParserOption {
	return func(p *Parser) {
		p.UseJSONNumber = true
	}
}

// WithoutClaimsValidation is an option to disable claims validation. This option should only be used if you exactly know
// what you are doing.
func WithoutClaimsValidation() ParserOption {
	return func(p *Parser) {
		p.SkipClaimsValidation = true
	}
}
