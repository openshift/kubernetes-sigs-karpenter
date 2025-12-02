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

package colorful

import (
	"math/rand"
)

// Uses the HSV color space to generate colors with similar S,V but distributed
// evenly along their Hue. This is fast but not always pretty.
// If you've got time to spare, use Lab (the non-fast below).
func FastHappyPalette(colorsCount int) (colors []Color) {
	colors = make([]Color, colorsCount)

	for i := 0; i < colorsCount; i++ {
		colors[i] = Hsv(float64(i)*(360.0/float64(colorsCount)), 0.8+rand.Float64()*0.2, 0.65+rand.Float64()*0.2)
	}
	return
}

func HappyPalette(colorsCount int) ([]Color, error) {
	pimpy := func(l, a, b float64) bool {
		_, c, _ := LabToHcl(l, a, b)
		return 0.3 <= c && 0.4 <= l && l <= 0.8
	}
	return SoftPaletteEx(colorsCount, SoftPaletteSettings{pimpy, 50, true})
}
