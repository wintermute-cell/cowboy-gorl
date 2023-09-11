package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    "math"
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

type number interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Max will return the maximum value between x and y
func Max[T number](x, y T) T {
    return T(math.Max(float64(x), float64(y)))
}

// Min will return the minimum value between x and y.
func Min[T number](x, y T) T {
    return T(math.Min(float64(x), float64(y)))
}

// Clamps x between lower_bound and upper_bound, both inclusive.
// (Clamp will return at least lower_bound and at most upper_bound)
func Clamp[T number](x, lower_bound, upper_bound T) T {
    v := Min[T](x, upper_bound)
    v = Max[T](v, lower_bound)
    return v
}
