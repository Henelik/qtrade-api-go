package v1

import "math"

// Round x to a specified number of decimal places
func Round(x float64, places int) float64 {
	factor := math.Pow(10, float64(places))

	return math.Round(x*factor) / factor
}
