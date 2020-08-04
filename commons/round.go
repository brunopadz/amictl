package commons

import "math"

func Round(x float64) float64 {
	return math.Round(x*100) / 100
}
