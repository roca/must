package must

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReturnElseLogFatal(t *testing.T) {

	// file exists
	content := ReturnElseLogFatal(ioutil.ReadFile, "example.txt").([]byte)
	if fmt.Sprintf("%T", content) != "[]byte" {
		t.Errorf("Incorrect Type for ioutil.ReadFile was returned. Expecting []byte got %T", content)
	}
}
