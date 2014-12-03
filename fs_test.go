package fs

import "fmt"
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

	fmt.Println(json)

	// if json == "{a:1, b:false}" {
	// 	t.Log("Passed")
	// } else {
	// 	t.Error("Failed: file not match")
	// }
}
