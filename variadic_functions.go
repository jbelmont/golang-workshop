package workshop

func sum(numbers ...float64) float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func average3(numbers ...float64) float64 {
	return sum(numbers...) / float64(len(numbers))
}

func variadicFunctions() {
	sumemup := func(numbers ...int) int {
		var sum int
		for _, v := range numbers {
			sum += v
		}
		return sum
	}

	numbers := []int{
		1, 2, 3, 4, 5,
	}
	assert(sumemup(numbers...) == 10)

	numbers2 := []float64{
		5.4, 2.3, 2, 9, 3.8, 1.9,
	}
	assert(average3(numbers2...) == 5.3)
}
