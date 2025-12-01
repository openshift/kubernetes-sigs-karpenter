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

package golines

import (
	"github.com/golangci/golines"

	"github.com/golangci/golangci-lint/v2/pkg/config"
)

const Name = "golines"

type Formatter struct {
	shortener *golines.Shortener
}

func New(settings *config.GoLinesSettings) *Formatter {
	options := golines.ShortenerConfig{}

	if settings != nil {
		options = golines.ShortenerConfig{
			MaxLen:           settings.MaxLen,
			TabLen:           settings.TabLen,
			KeepAnnotations:  false, // golines debug (not usable inside golangci-lint)
			ShortenComments:  settings.ShortenComments,
			ReformatTags:     settings.ReformatTags,
			IgnoreGenerated:  false, // handle globally
			DotFile:          "",    // golines debug (not usable inside golangci-lint)
			ChainSplitDots:   settings.ChainSplitDots,
			BaseFormatterCmd: "go fmt", // fake cmd
		}
	}

	return &Formatter{shortener: golines.NewShortener(options)}
}

func (*Formatter) Name() string {
	return Name
}

func (f *Formatter) Format(_ string, src []byte) ([]byte, error) {
	return f.shortener.Shorten(src)
}
