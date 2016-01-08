package utils

import (
	"testing"
)

func TestIntArrayToString(t *testing.T) {
	utils := GetUtils()
	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)
	str := utils.IntArrayToString(arr)
	if str != "1,2" {
		t.Error("Expected the string to be 1,2 got,", str)
	}
}g

