package err_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		//fmt.Println("Finally")
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something Wrong"))
	//os.Exit(-1)
}
