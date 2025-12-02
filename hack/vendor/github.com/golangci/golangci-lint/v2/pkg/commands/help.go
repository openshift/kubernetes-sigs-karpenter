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

package commands

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/v2/pkg/lint/lintersdb"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type helpOptions struct {
	JSON bool
}

type helpCommand struct {
	cmd *cobra.Command

	opts helpOptions

	dbManager *lintersdb.Manager

	log logutils.Log
}

func newHelpCommand(logger logutils.Log) *helpCommand {
	c := &helpCommand{log: logger}

	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Display extra help",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	lintersCmd := &cobra.Command{
		Use:               "linters",
		Short:             "Display help for linters.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              c.lintersExecute,
		PreRunE:           c.lintersPreRunE,
	}

	fsLinter := lintersCmd.Flags()
	fsLinter.SortFlags = false // sort them as they are defined here

	fsLinter.BoolVar(&c.opts.JSON, "json", false, color.GreenString("Display as JSON"))

	helpCmd.AddCommand(lintersCmd)

	formattersCmd := &cobra.Command{
		Use:               "formatters",
		Short:             "Display help for formatters.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              c.formattersExecute,
		PreRunE:           c.formattersPreRunE,
	}

	fsFormatter := formattersCmd.Flags()
	fsFormatter.SortFlags = false // sort them as they are defined here

	fsFormatter.BoolVar(&c.opts.JSON, "json", false, color.GreenString("Display as JSON"))

	helpCmd.AddCommand(formattersCmd)

	c.cmd = helpCmd

	return c
}

func formatDescription(desc string) string {
	desc = strings.TrimSpace(desc)

	if desc == "" {
		return desc
	}

	// If the linter description spans multiple lines, truncate everything following the first newline
	endFirstLine := strings.IndexRune(desc, '\n')
	if endFirstLine > 0 {
		desc = desc[:endFirstLine]
	}

	rawDesc := []rune(desc)

	r, _ := utf8.DecodeRuneInString(desc)
	rawDesc[0] = unicode.ToUpper(r)

	if rawDesc[len(rawDesc)-1] != '.' {
		rawDesc = append(rawDesc, '.')
	}

	return string(rawDesc)
}
