package test

import (
	"fmt"
	"testing"
)

func TestIncludes(t *testing.T) {
	num_list := []int{1, 2, 3, 4}
	var my_key int = 3

	for key, value := range num_list {
		fmt.Println(key, value)
		if value == my_key {
			fmt.Println("")
		}
	}
}
