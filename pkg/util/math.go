package util

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
	"time"
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

// DistributeInteger distributes the integer 'A' into buckets defined by their maximum capacities.
// It returns a slice of integers denoting the final distribution in each bucket.
func DistributeInteger(A int32, capacities []int32) []int32 {
	n := len(capacities)
	if n == 0 {
		fmt.Println("Error: The list of bucket capacities should not be empty.")
		return nil
	}

	// Create a slice to hold the final distribution in each bucket.
	distribution := make([]int32, n)

	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	for A > 0 {
		for i := 0; i < n; i++ {
			if A <= 0 {
				break
			}

			// Calculate the maximum we can put in the current bucket.
			maxAdd := int32(capacities[i]) - distribution[i]

			if maxAdd > 0 {
				// Decide a random amount to put in the current bucket (up to maxAdd or remaining A).
				add := int32(rand.Intn(int(math.Min(float64(maxAdd), float64(A)))) + 1)

				// Update the bucket and the remaining A.
				distribution[i] += add
				A -= add
			}
		}
	}

	return distribution
}
