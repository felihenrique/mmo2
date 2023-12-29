package ds

import "math"

func NormalizeVec(x float64, y float64) (float64, float64) {
	length := math.Sqrt(float64(x*x + y*y))

	// Check for division by zero to avoid NaN
	if length == 0 {
		return 0, 0
	}

	normalizedX := x / length
	normalizedY := y / length

	return normalizedX, normalizedY
}

func Distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}
