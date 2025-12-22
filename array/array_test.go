package array

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {

	str := ""

	arr, err := StringToUintSlice(str, ",")

	if err != nil {
		fmt.Println("=========err==========")
		fmt.Println("err:", err)
	}
	fmt.Println("arr:", arr)
}
