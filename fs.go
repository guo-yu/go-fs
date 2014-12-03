package fs

import "fmt"
import "os"

type JSON interface{}

func ReadFile(filename string) (fileContent string, err error) {

}

func WriteFile(filename string, fileContent string) error {

}

func ReadJSON(filename string) (json JSON, err error) {

}

func WriteJSON(filename string, jsonContent string) error {

}
