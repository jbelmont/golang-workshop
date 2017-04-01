package workshop

import (
	"errors"
	"log"
)

func add2Numbers(x, y int) int {
	return x + y
}

func printFormattedResponse(name string) string {
	return name + " hello there"
}

func wordSwap(word1, word2 string) (string, string) {
	return word2, word1
}

func nakedReturn(x, y, z *int) (x1, y1, z1 int) {
	*x++
	*y++
	*z++
	x1 = *x * 2
	y1 = *y * 3
	z1 = *z * 4
	return
}

func giveErrorIfEmpty(str string) (string, error) {
	if str == "" {
		error := errors.New("You must pass a string")
		return str, error
	}
	return str, nil
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

	lastName, firstName := wordSwap("marcel", "belmont")
	assert(lastName == "belmont" && firstName == "marcel")

	x, y, z := 1, 2, 3
	x2, y2, z2 := nakedReturn(&x, &y, &z)
	assert(x2 == 4 && y2 == 9 && z2 == 16)

	_, err := giveErrorIfEmpty("")
	assert(err.Error() == "You must pass a string")

	str, err := giveErrorIfEmpty("hello")
	if err != nil {
		log.Fatal(err)
	}
	assert(str == "hello")
}
