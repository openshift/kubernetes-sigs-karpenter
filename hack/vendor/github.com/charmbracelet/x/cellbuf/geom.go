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

package cellbuf

import (
	"image"
)

// Position represents an x, y position.
type Position = image.Point

// Pos is a shorthand for Position{X: x, Y: y}.
func Pos(x, y int) Position {
	return image.Pt(x, y)
}

// Rectange represents a rectangle.
type Rectangle = image.Rectangle

// Rect is a shorthand for Rectangle.
func Rect(x, y, w, h int) Rectangle {
	return image.Rect(x, y, x+w, y+h)
}
