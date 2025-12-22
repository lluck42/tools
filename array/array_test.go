package array

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {

	arr, err := StringToUintSlice("1,2-,3,4,5,6,7,8,9,10", ",")

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("arr:", arr)
}
