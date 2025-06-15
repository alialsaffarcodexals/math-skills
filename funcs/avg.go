package ms



func Avg(ai []int) float64 {
	s := 0
	for _, v := range ai {
		s += v
	}
	avg := float64(s) / float64(len(ai))
	return avg
}
