package workshop

import (
	"github.com/jbelmont/golang-workshop/doesExist"
	"github.com/jbelmont/golang-workshop/summer"
)

func packages() {
	movies := []string{
		"The Matrix", "Superman", "Platoon", "Footloose",
	}
	const movie = "The Fly"
	assert(doesExist.Exists(movies, movie) == true)

	numbers := []float64{
		1.9, 8.75, 1.3333, 988.3, 222.5, 333,
	}
	assert(summer.Sum(numbers) == 888)
}
