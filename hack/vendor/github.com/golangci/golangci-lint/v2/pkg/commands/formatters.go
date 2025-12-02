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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/lint/lintersdb"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type formattersHelp struct {
	Enabled  []formatterHelp
	Disabled []formatterHelp
}

type formattersOptions struct {
	config.LoaderOptions
	JSON bool
}

type formattersCommand struct {
	viper *viper.Viper
	cmd   *cobra.Command

	opts formattersOptions

	cfg *config.Config

	log logutils.Log

	dbManager *lintersdb.Manager
}

func newFormattersCommand(logger logutils.Log) *formattersCommand {
	c := &formattersCommand{
		viper: viper.New(),
		cfg:   config.NewDefault(),
		log:   logger,
	}

	formattersCmd := &cobra.Command{
		Use:               "formatters",
		Short:             "List current formatters configuration.",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              c.execute,
		PreRunE:           c.preRunE,
		SilenceUsage:      true,
	}

	fs := formattersCmd.Flags()
	fs.SortFlags = false // sort them as they are defined here

	setupConfigFileFlagSet(fs, &c.opts.LoaderOptions)

	setupFormattersFlagSet(c.viper, fs)

	fs.BoolVar(&c.opts.JSON, "json", false, color.GreenString("Display as JSON"))

	c.cmd = formattersCmd

	return c
}

func (c *formattersCommand) preRunE(cmd *cobra.Command, args []string) error {
	loader := config.NewFormattersLoader(c.log.Child(logutils.DebugKeyConfigReader), c.viper, cmd.Flags(), c.opts.LoaderOptions, c.cfg, args)

	err := loader.Load(config.LoadOptions{Validation: true})
	if err != nil {
		return fmt.Errorf("can't load config: %w", err)
	}

	dbManager, err := lintersdb.NewManager(c.log.Child(logutils.DebugKeyLintersDB), c.cfg,
		lintersdb.NewLinterBuilder(), lintersdb.NewPluginModuleBuilder(c.log), lintersdb.NewPluginGoBuilder(c.log))
	if err != nil {
		return err
	}

	c.dbManager = dbManager

	return nil
}

func (c *formattersCommand) execute(_ *cobra.Command, _ []string) error {
	enabledLintersMap, err := c.dbManager.GetEnabledLintersMap()
	if err != nil {
		return fmt.Errorf("can't get enabled formatters: %w", err)
	}

	var enabledFormatters []*linter.Config
	var disabledFormatters []*linter.Config

	for _, lc := range c.dbManager.GetAllSupportedLinterConfigs() {
		if lc.Internal {
			continue
		}

		if !goformatters.IsFormatter(lc.Name()) {
			continue
		}

		if enabledLintersMap[lc.Name()] == nil {
			disabledFormatters = append(disabledFormatters, lc)
		} else {
			enabledFormatters = append(enabledFormatters, lc)
		}
	}

	if c.opts.JSON {
		formatters := formattersHelp{}

		for _, lc := range enabledFormatters {
			formatters.Enabled = append(formatters.Enabled, newFormatterHelp(lc))
		}

		for _, lc := range disabledFormatters {
			formatters.Disabled = append(formatters.Disabled, newFormatterHelp(lc))
		}

		return json.NewEncoder(c.cmd.OutOrStdout()).Encode(formatters)
	}

	color.Green("Enabled by your configuration formatters:\n")
	printFormatters(enabledFormatters)

	color.Red("\nDisabled by your configuration formatters:\n")
	printFormatters(disabledFormatters)

	return nil
}
