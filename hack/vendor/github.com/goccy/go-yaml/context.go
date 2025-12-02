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

package yaml

import "context"

type (
	ctxMergeKey  struct{}
	ctxAnchorKey struct{}
)

func withMerge(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxMergeKey{}, true)
}

func isMerge(ctx context.Context) bool {
	v, ok := ctx.Value(ctxMergeKey{}).(bool)
	if !ok {
		return false
	}
	return v
}

func withAnchor(ctx context.Context, name string) context.Context {
	anchorMap := getAnchorMap(ctx)
	if anchorMap == nil {
		anchorMap = make(map[string]struct{})
	}
	anchorMap[name] = struct{}{}
	return context.WithValue(ctx, ctxAnchorKey{}, anchorMap)
}

func getAnchorMap(ctx context.Context) map[string]struct{} {
	v, ok := ctx.Value(ctxAnchorKey{}).(map[string]struct{})
	if !ok {
		return nil
	}
	return v
}
