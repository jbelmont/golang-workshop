package workshop

func slices() {
	var names = []string{"Jane", "Joe", "Margaret", "Michael", "Robin"}
	var sliceNames = make([]string, 5) // ["", "", "", "", ""]
	for i := 0; i < len(sliceNames); i++ {
		sliceNames[i] = names[i]
	}
	assert(sliceNames[2] == "Margaret")

	var newSlice = append(sliceNames, "John")
	assert(len(newSlice) == 6)

	var copyNewSlice = make([]string, len(newSlice))
	copy(copyNewSlice, newSlice)
	assert(len(newSlice) == len(copyNewSlice))

	subIndexSlice := copyNewSlice[:3]
	subIndexAgain := copyNewSlice[2:4]
	assert(subIndexSlice[2] == subIndexAgain[0])
}
