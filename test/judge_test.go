package test

import (
	"fmt"
	"testing"
)

func TestJudgeString(t *testing.T) {
	genderList := []string{"M", "F"}
	for _, value := range genderList {
		fmt.Println(value == "M")
	}
}
