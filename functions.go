package workshop

func add2Numbers(x, y int) int {
	return x + y
}

func printFormattedResponse(name string) string {
	return name + " hello there"
}

func functions() {
	adder := add2Numbers(5, 7)
	assert(adder == 12)

	formatResponse := printFormattedResponse("Jane")
	assert(formatResponse == "Jane hello there")

	summer := func(numbers []int) int {
		sum := 0
		for i := 0; i < len(numbers); i++ {
			sum += numbers[i]
		}
		return sum
	}

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	sumEmUp := summer(numbers)
	assert(sumEmUp == 28)
}
