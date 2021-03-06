package fs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Filemode os.FileMode = 0644

func ReadFile(filename string) (fileContent string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	fileContent = string(data)

	return
}

func WriteFile(filename string, fileContent string) error {
	return ioutil.WriteFile(filename, []byte(fileContent), Filemode)
}

func ReadJSON(filename string) (map[string]interface{}, error) {
	var ret map[string]interface{}
	var f interface{}

	fileContent, err := ReadFile(filename)

	if err != nil {
		return ret, err
	}

	err = json.Unmarshal([]byte(fileContent), &f)

	if err != nil {
		return ret, err
	}

	return f.(map[string]interface{}), nil
}

func WriteJSON(filename string, content interface{}) error {
	jsonString, err := json.Marshal(content)

	if err != nil {
		return err
	}

	ioutil.WriteFile(filename, jsonString, Filemode)

	return nil
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
