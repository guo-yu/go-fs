package fs

import "os"
import "path"
import "io/ioutil"
import "encoding/json"

func ReadFile(filename string) (fileContent string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(path.Join(cwd, filename))
	if err != nil {
		return "", err
	}

	fileContent = string(data)

	return
}

func WriteFile(filename string, fileContent string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(cwd, filename), []byte(fileContent), 0644)
}

func ReadJSON(filename string) (json string, err error) {
	return "", nil
}

func WriteJSON(filename string, jsonContent string) error {
	return nil
}
