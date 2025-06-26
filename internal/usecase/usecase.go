package sum

func CalculateSum(values []float64) (sum float64) {
	for _, v := range values {
		sum += v
	}
	return
}
