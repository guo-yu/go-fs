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
