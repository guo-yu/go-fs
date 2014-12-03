package fs

import "os"
import "path"
import "io/ioutil"
import "encoding/json"

const FILEMODE = 0644

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

	return ioutil.WriteFile(path.Join(cwd, filename), []byte(fileContent), FILEMODE)
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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	jsonString, err2 := json.Marshal(content)

	if err2 != nil {
		return err2
	}

	ioutil.WriteFile(path.Join(cwd, filename), jsonString, FILEMODE)

	return nil
}
