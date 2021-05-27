package must

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestReturnElseLogFatal(t *testing.T) {
	// If file exists no fatal error raised
	content := ReturnElseLogFatal(ioutil.ReadFile, "example.txt").([]byte)
	t.Log(reflect.TypeOf(content))
	if fmt.Sprintf("%T", content) != "[]uint8" {
		t.Errorf("Incorrect Type for ioutil.ReadFile was returned. Expecting []byte got %T", content)
	}
}

func TestReturnElseLogFatal02(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
            t.Errorf("Did not recover from function failure")
        }
	}()
	
	//If file does exists a panic is raised
	_ = ReturnElseLogFatal(ioutil.ReadFile, "exampleNonExisting.txt").([]byte)
}

