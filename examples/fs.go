package main

import "os"
import "path"
import "fmt"
import ".."

var cwd, err = os.Getwd()

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println(filename)

	// Create a tmp file
	err := fs.WriteFile(path.Join(cwd, "tmp.txt"), "demo text")
	check(err)

	// Read from tmp file we just created
	fileString, err := fs.ReadFile(path.Join(cwd, "tmp.txt"))
	check(err)

	// => `demo text`
	fmt.Println(fileString)

	data := map[string]string{
		"a": "123",
		"b": "456",
	}

	// Create JSON file from a map
	err = fs.WriteJSON(path.Join(cwd, "tmp.json"), data)
	check(err)

	// Read JSON file from tmp json file
	json, err := fs.ReadJSON(path.Join(cwd, "tmp.json"))

	fmt.Println(json)
}
