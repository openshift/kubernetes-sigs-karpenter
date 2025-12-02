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
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/lint/lintersdb"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type linterHelp struct {
	Name        string   `json:"name"`
	Desc        string   `json:"description"`
	Groups      []string `json:"groups"`
	Fast        bool     `json:"fast"`
	AutoFix     bool     `json:"autoFix"`
	Deprecated  bool     `json:"deprecated"`
	Since       string   `json:"since"`
	OriginalURL string   `json:"originalURL,omitempty"`
}

func newLinterHelp(lc *linter.Config) linterHelp {
	groups := []string{config.GroupAll}

	if !lc.IsSlowLinter() {
		groups = append(groups, config.GroupFast)
	}

	return linterHelp{
		Name:        lc.Name(),
		Desc:        formatDescription(lc.Linter.Desc()),
		Groups:      slices.Concat(groups, slices.Collect(maps.Keys(lc.Groups))),
		Fast:        !lc.IsSlowLinter(),
		AutoFix:     lc.CanAutoFix,
		Deprecated:  lc.IsDeprecated(),
		Since:       lc.Since,
		OriginalURL: lc.OriginalURL,
	}
}

func (c *helpCommand) lintersPreRunE(_ *cobra.Command, _ []string) error {
	// The command doesn't depend on the real configuration.
	dbManager, err := lintersdb.NewManager(c.log.Child(logutils.DebugKeyLintersDB), config.NewDefault(), lintersdb.NewLinterBuilder())
	if err != nil {
		return err
	}

	c.dbManager = dbManager

	return nil
}

func (c *helpCommand) lintersExecute(_ *cobra.Command, _ []string) error {
	if c.opts.JSON {
		return c.lintersPrintJSON()
	}

	c.lintersPrint()

	return nil
}

func (c *helpCommand) lintersPrintJSON() error {
	var linters []linterHelp

	for _, lc := range c.dbManager.GetAllSupportedLinterConfigs() {
		if lc.Internal {
			continue
		}

		if goformatters.IsFormatter(lc.Name()) {
			continue
		}

		linters = append(linters, newLinterHelp(lc))
	}

	return json.NewEncoder(c.cmd.OutOrStdout()).Encode(linters)
}

func (c *helpCommand) lintersPrint() {
	var enabledLCs, disabledLCs []*linter.Config
	for _, lc := range c.dbManager.GetAllSupportedLinterConfigs() {
		if lc.Internal {
			continue
		}

		if goformatters.IsFormatter(lc.Name()) {
			continue
		}

		if lc.FromGroup(config.GroupStandard) {
			enabledLCs = append(enabledLCs, lc)
		} else {
			disabledLCs = append(disabledLCs, lc)
		}
	}

	color.Green("Enabled by default linters:\n")
	printLinters(enabledLCs)

	color.Red("\nDisabled by default linters:\n")
	printLinters(disabledLCs)
}

func printLinters(lcs []*linter.Config) {
	slices.SortFunc(lcs, func(a, b *linter.Config) int {
		if a.IsDeprecated() && b.IsDeprecated() {
			return strings.Compare(a.Name(), b.Name())
		}

		if a.IsDeprecated() {
			return 1
		}

		if b.IsDeprecated() {
			return -1
		}

		return strings.Compare(a.Name(), b.Name())
	})

	for _, lc := range lcs {
		desc := formatDescription(lc.Linter.Desc())

		deprecatedMark := ""
		if lc.IsDeprecated() {
			deprecatedMark = " [" + color.RedString("deprecated") + "]"
		}

		var capabilities []string
		if !lc.IsSlowLinter() {
			capabilities = append(capabilities, color.BlueString("fast"))
		}
		if lc.CanAutoFix {
			capabilities = append(capabilities, color.GreenString("auto-fix"))
		}

		var capability string
		if capabilities != nil {
			capability = " [" + strings.Join(capabilities, ", ") + "]"
		}

		_, _ = fmt.Fprintf(logutils.StdOut, "%s%s: %s%s\n",
			color.YellowString(lc.Name()), deprecatedMark, desc, capability)
	}
}
