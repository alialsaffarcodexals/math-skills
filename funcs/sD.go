package ms

import "math"

func SD(Var int) int {
	t := math.Sqrt(float64(Var))
	return int(math.Round(t))
}
