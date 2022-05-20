package interop

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Fatal(err error, t *testing.T) {
	if err != nil {
		fmt.Println("detail error msg")
		fmt.Println(err)
		t.Fatal(err)
	}
}
func ReadFile(filename string, t *testing.T) []byte {
	content, err := ioutil.ReadFile(filename)
	fmt.Println(filename)
	fmt.Println("File content")
	fmt.Println(string(content))
	Fatal(err, t)
	return con