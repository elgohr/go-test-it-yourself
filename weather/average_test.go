package weather_test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"workshop/weather"
)

func TestCalculateAverage(t *testing.T) {
	require.Equal(t, 1, weather.CalculateAverage([]int{1}))
	require.Equal(t, 2, weather.CalculateAverage([]int{1,2,3}))
	require.Equal(t, 0, weather.CalculateAverage([]int{-1,1}))
	require.Equal(t, 0, weather.CalculateAverage([]int{}))
}
