package fs

import "os"
import "path"
import "testing"

var cwd, err = os.Getwd()

func TestFsReadFile(t *testing.T) {
	str, _ := ReadFile(path.Join(cwd, "examples/test.txt"))

	if str == "123456" {
		t.Log("Passed!")
	} else {
		t.Error("Failed!")
	}
}

func TestFsWriteFile(t *testing.T) {
	err := WriteFile(path.Join(cwd, "examples/tmp.txt"), "tmp file content")

	if err != nil {
		t.Error(err)
	} else {
		t.Log("Passed!")
	}
}

func TestFsReadJSON(t *testing.T) {
	json, err := ReadJSON(path.Join(cwd, "examples/test.json"))

	if err != nil {
		t.Error(err)
		return
	}

	if json["b"] == false {
		t.Log("Passed")
	} else {
		t.Error("Did not match")
	}
}

func TestFsWriteJSON(t *testing.T) {
	demoJson := map[string]string{
		"a": "abcde",
		"b": "cswsd",
	}

	err := WriteJSON(path.Join(cwd, "examples/tmp.json"), demoJson)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("Passed!")
	}
}

func TestFsExists(t *testing.T) {
	fsGoExists := Exists("./fs.go")

	if fsGoExists {
		t.Log("Passed")
	} else {
		t.Error("Failed")
	}
}
