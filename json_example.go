package workshop

import (
	"encoding/json"
	"fmt"
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

	slice := []string{"jack", "be", "nimble"}
	typeSlice, err := json.Marshal(slice)

	var parsed []interface{}
	err2 := json.Unmarshal(typeSlice, &parsed)
	if err2 != nil {
		panic(err2)
	}
	var expectedSlice = []string{"jack", "be", "nimble"}
	var sliceCompare = true
	for idx, val := range parsed {
		if val != expectedSlice[idx] {
			sliceCompare = false
		}
	}
	assert(sliceCompare == true)

	data := []byte(`
{
	"names": ["James", "Jane", "Chen"],
	"hash": {
		"secret": "12345abcdef"
	}
}
	`)
	var parsed2 map[string]interface{}
	err3 := json.Unmarshal(data, &parsed2)
	if err3 != nil {
		panic(err3)
	}
	for _, value := range parsed2 {
		switch value.(type) {
		case []interface{}:
			expectedNestedSlice := map[string]interface{}{
				"names": []string{"James", "Jane", "Chen"},
				"hash":  map[string]interface{}{"secret": "12345abcdef"},
			}
			_, ok := value.([]interface{})
			_, ok2 := expectedNestedSlice["names"].([]interface{})
			assert(ok != ok2)
		case map[string]interface{}:
			val := value.(map[string]interface{})["secret"]
			assert(val == "12345abcdef")
		default:
			fmt.Println("")
		}
	}

	type Slices struct {
		Names []string               `json:"names"`
		Hash  map[string]interface{} `json:"hash"`
	}
	var knownJSON Slices
	err4 := json.Unmarshal(data, &knownJSON)
	if err4 != nil {
		panic(err4)
	}

	namesSlice := knownJSON.Names
	var namesSliceExpected = []string{"James", "Jane", "Chen"}
	assertSlices := true
	for idx, vals := range namesSlice {
		if vals != namesSliceExpected[idx] {
			assertSlices = false
		}
	}
	assert(assertSlices == true)

	actualSecret := knownJSON.Hash["secret"]
	secretHash := map[string]string{
		"secret": "12345abcdef",
	}
	assert(actualSecret == secretHash["secret"])
}
