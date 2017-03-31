package practice

import (
	"math"
	"strconv"
)

func averageEmUp(numbers []float64) float64 {
	return average(numbers)
}

func average(numbers []float64) float64 {
	sum := 0.0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / float64(len(numbers))
}

func format2DecimalPlaces(num float64) string {
	str := strconv.FormatFloat(num, 'f', 2, 64)
	return str
}

func standardDeviation(numbers []float64) func() string {
	mean := average(numbers)
	var summation float64
	return func() string {
		for i := 0; i < len(numbers); i++ {
			summation += math.Pow(math.Abs((numbers[i] - mean)), 2)
		}
		format2Dec := format2DecimalPlaces(math.Sqrt(summation / float64(len(numbers))))
		return format2Dec
	}
}
