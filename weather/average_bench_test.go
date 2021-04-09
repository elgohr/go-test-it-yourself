package weather_test

import (
	"testing"
	"workshop/weather"
)

func BenchmarkCalculateAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		weather.CalculateAverage([]int{1,2,3})
	}
}

func BenchmarkCalculateStandardLibAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlternativeCalculateAverage([]int{1,2,3})
	}
}

func AlternativeCalculateAverage(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum / len(values)
}
