package ms

import (
	"math"
	"sort"
)

func Med(ai []int) int {
	sort.Ints(ai)
	l := len(ai)
	if l%2 == 0 {
		a := float64(ai[l/2])
		b := float64(ai[(l/2)-1])
		return int(math.Round((a + b) / 2))
	}
	return ai[l/2]
}
