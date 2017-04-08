package workshop

func slices() {
	var names = []string{"Jane", "Joe", "Margaret", "Michael", "Robin"}
	var sliceNames = make([]string, 5) // ["", "", "", "", ""]
	for i := 0; i < len(sliceNames); i++ {
		sliceNames[i] = names[i]
	}
	assert(sliceNames[2] == "Robin")

	var newSlice = append(sliceNames, "John")
	assert(len(newSlice) == 5)

	var copyNewSlice = make([]string, len(newSlice))
	copy(copyNewSlice, newSlice)
	assert(len(newSlice) == 0)

	subIndexSlice := copyNewSlice[:3]
	subIndexAgain := copyNewSlice[2:4]
	assert(subIndexSlice[2] == subIndexAgain[2])
}
