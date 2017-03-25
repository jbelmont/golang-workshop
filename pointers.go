package workshop

type avg []float64

func (nums *avg) average() float64 {
	var sum float64
	for _, num := range *nums {
		sum += num
	}
	return sum / float64(len((*nums)))
}

func pointers() {
	i, j := 3, 4
	var a = &i
	var b = &j
	assert(i == *a)
	assert(j == *b)

	*a = 10
	assert(i == *a)

	var num *int
	assert(num == nil)
	var numbers = avg{2.3, 3.5, 7.5, 2.5, 7.8}
	assert(numbers.average() == 4.720000000000001)
}
