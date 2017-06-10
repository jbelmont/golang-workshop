package workshop

import (
	"errors"
	"log"
	"math"
	"strconv"
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

func hypotenuse(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func summation(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func format2DecimalPlaces(num float64) string {
	str := strconv.FormatFloat(num, 'f', 2, 64)
	return str
}

func distanceOfLine(x1, y1, x2, y2 float64) func() string {
	var distance float64
	return func() string {
		distance = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
		return format2DecimalPlaces(distance)
	}
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

	hyp := hypotenuse(3, 4)
	assert(hyp == 5)

	sum1 := summation(1, 2, 3)
	assert(sum1 == 6)

	sum2 := summation()
	assert(sum2 == 0)

	distance := distanceOfLine(3, 4, 5, 6)
	assert(distance() == "2.83")
}
