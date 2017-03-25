package workshop

import (
	"encoding/json"
	"strconv"
)

func jsonExamples() {
	// Marshalling means converting to JSON in Golang
	typeBool, err := json.Marshal(true)
	if err != nil {
		panic(err)
	}
	convertTypeBool, err := strconv.ParseBool(string(typeBool))
	if err != nil {
		panic(err)
	}
	assert(convertTypeBool == true)

	// json.Marshal encodes data type and takes in interface{} which is any go type and returns a []byte
	typeNumber, err := json.Marshal(10)
	convertTypeNumber, err := strconv.ParseInt(string(typeNumber), 10, 64)
	if err != nil {
		panic(err)
	}
	assert(convertTypeNumber == 10)

	typeString, err := json.Marshal("Program")
	if err != nil {
		panic(err)
	}
	convertTypeString := string(typeString[:])
	convertExpected := "Program"
	assert(convertTypeString != convertExpected)
}
