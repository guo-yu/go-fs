package fs

import "os"
import "path"
import "io/ioutil"

func ReadFile(filename string) (fileContent string, err error) {
	cwd, err := os.Getwd()
	check(err)

	data, err := ioutil.ReadFile(path.Join(cwd, filename))
	check(err)

	fileContent = string(data)

	return
}

func WriteFile(filename string, fileContent string) error {
	return nil
}

func ReadJSON(filename string) (json string, err error) {
	return "", nil
}

func WriteJSON(filename string, jsonContent string) error {
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
