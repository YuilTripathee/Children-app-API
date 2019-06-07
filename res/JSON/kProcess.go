package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// for performing linear search
func findAlphabet(char1 string, result []map[string]interface{}) map[string]interface{} {
	for i := 0; i < len(result); i++ {
		if result[i]["l_case"] == char1 {
			return result[i]
		}
	}
	return nil
}

// BinarySearch (string character, array of objects(LOL) ) // used for performing binary search
func BinarySearch(char1 string, result []map[string]interface{}) map[string]interface{} {
	mid := len(result) / 2

	switch {
	case len(result) == 0:
		return nil // since the supplied array is an empty array
	case result[mid]["l_case"].(string) > char1:
		return BinarySearch(char1, result[:mid])
	case result[mid]["l_case"].(string) < char1:
		return BinarySearch(char1, result[mid+1:])
	default:
		return result[mid]
	}
}

func main() {
	fileAddr := "data/alphabets_raw.json"

	// open our json file
	jsonFile, err := os.Open(fileAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sucessfully opened the file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result[25])

	fmt.Println(findAlphabet("a", result))
	fmt.Println(BinarySearch("m", result))

	// fmt.Println(AMD)
}

/*
	Refereces
	1. https://tutorialedge.net/golang/parsing-json-with-golang/
	2. https://golangcode.com/json-encode-an-array-of-objects/ (for next project)
	3. https://www.geeksforgeeks.org/binary-search-a-string/

	To-do (later on-demand):
	1. Conversion of array of Interfaces to Array of Structures (in GOLANG)
*/
