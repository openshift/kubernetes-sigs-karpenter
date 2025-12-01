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

package goutil

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ldez/grignotin/goenv"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type EnvKey string

type Env struct {
	vars map[string]string
	log  logutils.Log
}

func NewEnv(log logutils.Log) *Env {
	return &Env{
		vars: map[string]string{},
		log:  log,
	}
}

func (e Env) Discover(ctx context.Context) error {
	startedAt := time.Now()

	var err error
	e.vars, err = goenv.Get(ctx, goenv.GOCACHE, goenv.GOROOT)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	e.log.Infof("Read go env for %s: %#v", time.Since(startedAt), e.vars)

	return nil
}

func (e Env) Get(k EnvKey) string {
	envValue := os.Getenv(string(k))
	if envValue != "" {
		return envValue
	}

	return e.vars[string(k)]
}
