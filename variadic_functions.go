package workshop

func variadicFunctions() {
	sumemup := func(numbers ...int) int {
		var sum int
		for _, v := range numbers {
			sum += v
		}
		return sum
	}

	numbers := []int{
		1,2,3,4,5,
	}
	assert(sumemup(numbers...) == 10)
}