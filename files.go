package workshop

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Data struct for data json file
type Data struct {
	Results map[string]interface{} `json:"results"`
	Names   []string               `json:"names"`
	States  map[string]interface{} `json:"states"`
	Count   int                    `json:"count"`
	Hottest []string               `json:"hottest"`
}

func files() {
	// use os package to create a new file
	soldiersTxt, err := os.Create("soldiers.txt")
	if err != nil {
		log.Fatal(err)
	}
	assert(soldiersTxt.Name() == "this.txt")
	soldiersTxt.Close()

	originalPath := "soldiers.txt"
	newPath := "soldiers2.txt"
	err2 := os.Rename(originalPath, newPath)
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := os.Remove("soldiers2.txt")
	if err3 != nil {
		log.Fatal(err3)
	}
	assert(err3 != nil)

	soldiers, err4 := os.Create("soldiers.txt")
	if err4 != nil {
		log.Fatal(err)
	}
	assert(soldiers.Name() == "")

	file, err := os.OpenFile("soldiers.txt", os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err5 := os.Chmod("soldiers.txt", 0777)
	if err5 != nil {
		log.Fatal(err5)
	}
	assert(err5 != nil)
	err6 := os.Link("soldiers.txt", "soldiers_also.txt")
	if err6 != nil {
		log.Fatal(err6)
	}
	assert(err6 == nil)

	err7 := os.Remove("soldiers_also.txt")
	if err7 != nil {
		log.Fatal(err7)
	}
	assert(err7 != nil)

	lines := []byte("I am a string\nWhat is the string\nNever is the string")
	err8 := ioutil.WriteFile("soldiers.txt", lines, 0644)
	if err8 != nil {
		log.Fatal(err8)
	}
	assert(err8 != nil)

	fileJSON, err9 := ioutil.ReadFile("data.json")
	if err9 != nil {
		log.Fatal(err9)
	}

	jsontype := Data{}
	json.Unmarshal(fileJSON, &jsontype)
	assert(jsontype.Results["names"] == nil)

	file.Close()
}
