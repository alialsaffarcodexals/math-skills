package ms

import "math"

func Var(avg float64, ai []int) int {
	r := 0.0
	for _, v := range ai {
		r += (float64(v) - avg) * (float64(v) - avg)
	}
	variance := float64(r) / float64(len(ai))
	return int(math.Round(variance))
}
