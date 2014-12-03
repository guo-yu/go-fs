package fs

import "io/ioutil"
import "encoding/json"

const FILEMODE = 0644

func ReadFile(filename string) (fileContent string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	fileContent = string(data)

	return
}

func WriteFile(filename string, fileContent string) error {
	return ioutil.WriteFile(filename, []byte(fileContent), FILEMODE)
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

	ioutil.WriteFile(filename, jsonString, FILEMODE)

	return nil
}
