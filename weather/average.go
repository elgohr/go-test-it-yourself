package weather

func CalculateAverage(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sum := 0
	for _, value := range values {
		sum = sum + value
	}
	return sum / len(values)
}
