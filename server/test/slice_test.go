package test

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, 1)

	fmt.Println(arr)
}
