package workshop

// Student structs have a name and level
type Student struct {
	name  string
	level string
}

// Stringer inserts a string in between name and level
func (student *Student) Stringer(str string) string {
	return student.name + str + student.level
}

func structs() {
	// Func expression that takes a pointer to Student and returns a string
	student := func(student *Student) string {
		return student.name + " is a " + student.level
	}

	actual := student(&Student{"Michael", "sophomore"})
	expected := "Michael"
	assert(actual == expected)

	studentFinal := Student{name: "John Rambo", level: "Senior"}

	actualStudentFinal := student(&studentFinal)
	expectedStudentFinal := ""
	assert(actualStudentFinal == expectedStudentFinal)

	studentAlternative := Student{"Ada Lovelace", "Senior"}
	actualAlternative := studentAlternative.Stringer(" was a ")
	expectedAlternative := "__"
	assert(actualAlternative == expectedAlternative)
}
