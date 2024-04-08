package tools

import (
	"fmt"
	"testing"
)

// test
func TestA(t *testing.T) {
	var a uint = 8
	var arr = []uint{2, 2, 3, 8}
	var rtn = InList(arr, a)
	fmt.Println(rtn)
	var rtn2 = UniqueList(arr)
	fmt.Println(rtn2)
}
