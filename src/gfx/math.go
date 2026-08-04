package gfx

import (
	"math"
)

const Tau = math.Pi * 2

// Cosine function that uses turns instead of radians.
func CosT(v float32) float32 {
	return float32(math.Cos(float64(v * Tau)))
}

// Cosine function that uses turns instead of radians.
func SinT(v float32) float32 {
	return float32(math.Sin(float64(v * Tau)))
}
