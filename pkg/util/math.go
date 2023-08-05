package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Vector2Clamp restricts a vector within the limits specified by min and max vectors.
func Vector2Clamp(input, min, max rl.Vector2) rl.Vector2 {
	if input.X < min.X {
		input.X = min.X
	} else if input.X > max.X {
		input.X = max.X
	}
	if input.Y < min.Y {
		input.Y = min.Y
	} else if input.Y > max.Y {
		input.Y = max.Y
	}
	return input
}
