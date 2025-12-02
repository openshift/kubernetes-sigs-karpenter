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

package command

import (
	"fmt"
	"io"
	"strings"

	"github.com/onsi/ginkgo/v2/formatter"
	"github.com/onsi/ginkgo/v2/types"
)

type Command struct {
	Name          string
	Flags         types.GinkgoFlagSet
	Usage         string
	ShortDoc      string
	Documentation string
	DocLink       string
	Command       func(args []string, additionalArgs []string)
}

func (c Command) Run(args []string, additionalArgs []string) {
	args, err := c.Flags.Parse(args)
	if err != nil {
		AbortWithUsage(err.Error())
	}
	for _, arg := range args {
		if len(arg) > 1 && strings.HasPrefix(arg, "-") {
			AbortWith(types.GinkgoErrors.FlagAfterPositionalParameter().Error())
		}
	}
	c.Command(args, additionalArgs)
}

func (c Command) EmitUsage(writer io.Writer) {
	fmt.Fprintln(writer, formatter.F("{{bold}}"+c.Usage+"{{/}}"))
	fmt.Fprintln(writer, formatter.F("{{gray}}%s{{/}}", strings.Repeat("-", len(c.Usage))))
	if c.ShortDoc != "" {
		fmt.Fprintln(writer, formatter.Fiw(0, formatter.COLS, c.ShortDoc))
		fmt.Fprintln(writer, "")
	}
	if c.Documentation != "" {
		fmt.Fprintln(writer, formatter.Fiw(0, formatter.COLS, c.Documentation))
		fmt.Fprintln(writer, "")
	}
	if c.DocLink != "" {
		fmt.Fprintln(writer, formatter.Fi(0, "{{bold}}Learn more at:{{/}} {{cyan}}{{underline}}http://onsi.github.io/ginkgo/#%s{{/}}", c.DocLink))
		fmt.Fprintln(writer, "")
	}
	flagUsage := c.Flags.Usage()
	if flagUsage != "" {
		fmt.Fprintf(writer, formatter.F(flagUsage))
	}
}
