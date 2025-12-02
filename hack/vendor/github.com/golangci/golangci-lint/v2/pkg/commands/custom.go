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
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/v2/pkg/commands/internal"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

const envKeepTempFiles = "CUSTOM_GCL_KEEP_TEMP_FILES"

type customCommand struct {
	cmd *cobra.Command

	cfg *internal.Configuration

	log logutils.Log
}

func newCustomCommand(logger logutils.Log) *customCommand {
	c := &customCommand{log: logger}

	customCmd := &cobra.Command{
		Use:          "custom",
		Short:        "Build a version of golangci-lint with custom linters.",
		Args:         cobra.NoArgs,
		PreRunE:      c.preRunE,
		RunE:         c.runE,
		SilenceUsage: true,
	}

	c.cmd = customCmd

	return c
}

func (c *customCommand) preRunE(_ *cobra.Command, _ []string) error {
	cfg, err := internal.LoadConfiguration()
	if err != nil {
		return err
	}

	err = cfg.Validate()
	if err != nil {
		return err
	}

	c.cfg = cfg

	return nil
}

func (c *customCommand) runE(cmd *cobra.Command, _ []string) error {
	tmp, err := os.MkdirTemp(os.TempDir(), "custom-gcl")
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}

	defer func() {
		if os.Getenv(envKeepTempFiles) != "" {
			log.Printf("WARN: The env var %s has been detected: the temporary directory is preserved: %s", envKeepTempFiles, tmp)

			return
		}

		_ = os.RemoveAll(tmp)
	}()

	err = internal.NewBuilder(c.log, c.cfg, tmp).Build(cmd.Context())
	if err != nil {
		return fmt.Errorf("build process: %w", err)
	}

	return nil
}
