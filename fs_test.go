package fs

import "testing"

func TestFsReadFile(t *testing.T) {
	str, _ := ReadFile("examples/test.txt")

	if str == "123456" {
		t.Log("Passed!")
	} else {
		t.Error("Failed!")
	}
}

func TestFsWriteFile(t *testing.T) {
	err := WriteFile("examples/tmp.txt", "tmp file content")

	if err != nil {
		t.Error(err)
	} else {
		t.Log("Passed!")
	}
}

func TestFsReadJSON(t *testing.T) {
	json, err := ReadJSON("examples/test.json")

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

	err := WriteJSON("examples/tmp.json", demoJson)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("Passed!")
	}
}
