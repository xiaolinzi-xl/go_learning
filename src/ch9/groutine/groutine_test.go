package groutine_test

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		//time.Sleep(time.Millisecond * 10)
		//go func() {
		//	fmt.Println(i)
		//}()
	}
}
