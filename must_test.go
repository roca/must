package must

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestReturnElseLogFatal(t *testing.T) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
				t.Log(err)
			}
		}
	}()

	// If file exists no fatal error raised
	content := ReturnElseLogFatal(ioutil.ReadFile, "example.txt").([]byte)
	t.Log(reflect.TypeOf(content))
	if fmt.Sprintf("%T", content) != "[]uint8" {
		t.Errorf("Incorrect Type for ioutil.ReadFile was returned. Expecting []byte got %T", content)
	}

	// // If file does exists a fatal error is raised
	// content2 := ReturnElseLogFatal(ioutil.ReadFile, "exampleNonExisting.txt").([]byte)
	// if err == nil {
	// 	t.Errorf("A fatal error should have been raised %T", content2)
	// }

}
