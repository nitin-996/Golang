package main

import (
	
	"testing"
)

func Test_checkEven(t *testing.T) {
    i := 9
    expected := "evennn" // This should cause a failure
    result := checkEven(i)
    if result != expected {
        t.Errorf("Expected %v but got %v", expected, result)
    }
}
