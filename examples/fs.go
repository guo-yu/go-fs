package main

import "fmt"
import ".."

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create a tmp file
	err := fs.WriteFile("./tmp.txt", "demo text")
	check(err)

	// Read from tmp file we just created
	fileString, err := fs.ReadFile("./tmp.txt")
	check(err)

	// => `demo text`
	fmt.Println(fileString)

	data := map[string]string{
		"a": "123",
		"b": "456",
	}

	// Create JSON file from a map
	err = fs.WriteJSON("./tmp.json", data)
	check(err)

	// Read JSON file from tmp json file
	json, err := fs.ReadJSON("./tmp.json")

	fmt.Println(json)
}
