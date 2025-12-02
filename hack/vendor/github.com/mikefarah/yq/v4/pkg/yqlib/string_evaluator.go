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

package yqlib

import (
	"bufio"
	"bytes"
	"container/list"
	"strings"
)

type StringEvaluator interface {
	Evaluate(expression string, input string, encoder Encoder, decoder Decoder) (string, error)
	EvaluateAll(expression string, input string, encoder Encoder, decoder Decoder) (string, error)
}

type stringEvaluator struct {
	treeNavigator DataTreeNavigator
}

func NewStringEvaluator() StringEvaluator {
	return &stringEvaluator{
		treeNavigator: NewDataTreeNavigator(),
	}
}

func (s *stringEvaluator) EvaluateAll(expression string, input string, encoder Encoder, decoder Decoder) (string, error) {
	reader := bufio.NewReader(strings.NewReader(input))
	var documents *list.List
	var results *list.List
	var err error

	if documents, err = ReadDocuments(reader, decoder); err != nil {
		return "", err
	}

	evaluator := NewAllAtOnceEvaluator()
	if results, err = evaluator.EvaluateCandidateNodes(expression, documents); err != nil {
		return "", err
	}

	out := new(bytes.Buffer)
	printer := NewPrinter(encoder, NewSinglePrinterWriter(out))
	if err := printer.PrintResults(results); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (s *stringEvaluator) Evaluate(expression string, input string, encoder Encoder, decoder Decoder) (string, error) {

	// Use bytes.Buffer for output of string
	out := new(bytes.Buffer)
	printer := NewPrinter(encoder, NewSinglePrinterWriter(out))

	InitExpressionParser()
	node, err := ExpressionParser.ParseExpression(expression)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(strings.NewReader(input))
	evaluator := NewStreamEvaluator()
	if _, err := evaluator.Evaluate("", reader, node, printer, decoder); err != nil {
		return "", err
	}
	return out.String(), nil
}
