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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/v2/internal/cache"
	"github.com/golangci/golangci-lint/v2/pkg/fsutils"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type cacheCommand struct {
	cmd *cobra.Command
}

func newCacheCommand() *cacheCommand {
	c := &cacheCommand{}

	cacheCmd := &cobra.Command{
		Use:   "cache",
		Short: "Cache control and information.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cacheCmd.AddCommand(
		&cobra.Command{
			Use:               "clean",
			Short:             "Clean cache",
			Args:              cobra.NoArgs,
			ValidArgsFunction: cobra.NoFileCompletions,
			RunE:              c.executeClean,
		},
		&cobra.Command{
			Use:               "status",
			Short:             "Show cache status",
			Args:              cobra.NoArgs,
			ValidArgsFunction: cobra.NoFileCompletions,
			Run:               c.executeStatus,
		},
	)

	c.cmd = cacheCmd

	return c
}

func (*cacheCommand) executeClean(_ *cobra.Command, _ []string) error {
	cacheDir := cache.DefaultDir()

	if err := os.RemoveAll(cacheDir); err != nil {
		return fmt.Errorf("failed to remove dir %s: %w", cacheDir, err)
	}

	return nil
}

func (*cacheCommand) executeStatus(_ *cobra.Command, _ []string) {
	cacheDir := cache.DefaultDir()

	_, _ = fmt.Fprintf(logutils.StdOut, "Dir: %s\n", cacheDir)

	cacheSizeBytes, err := dirSizeBytes(cacheDir)
	if err == nil {
		_, _ = fmt.Fprintf(logutils.StdOut, "Size: %s\n", fsutils.PrettifyBytesCount(cacheSizeBytes))
	}
}

func dirSizeBytes(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
