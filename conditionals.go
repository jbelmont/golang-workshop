package workshop

import (
	"math"
	"strconv"
	"strings"
)

func conditionals() {
	var str string
	flag := true

	if flag {
		str = "Righteous Path"
	}
	assert(str == "Righteous Path")

	var value int
	if value < 10 {
		value = 10
	}
	assert(value == 10)

	value2 := 7
	if value2 > 10 {
		value2++
	} else {
		value2--
	}
	assert(value2 == 6)

	const movie = "Rocky"
	var score int
	if movie == "Rambo" {
		score = 9
	} else if movie == "Rocky" {
		score = 10
	} else {
		score = 6
	}
	assert(score == 10)

	const food = "Pizza"
	var drink string
	switch food {
	case "Fries":
		drink = "Coke"
	case "Sandwich":
		drink = "Sprite"
	default:
		drink = "Milkshake"
	}
	assert(drink == "Milkshake")

	rem := math.Mod(6, 5) == 1
	assert(rem == true)

	num := "5"
	otherNum, err := strconv.Atoi(num)
	if err == nil {
		otherNum--
	}
	assert(otherNum == 4)

	strVal := "The Knowledge"

	// conditional can do an assignment and then have conditional
	// strings.Contains returns a boolean
	if ok := strings.Contains(strVal, "Knowledge"); ok {
		strVal += " is there"
	}
	assert(strVal == "The Knowledge is there")
}
