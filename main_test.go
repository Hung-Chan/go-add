package main

import "testing"

func TestAdd(t *testing.T) {
	testNum1 := 2
	testNum2 := 1

	testTotal := add(testNum1, testNum2)

	t.Logf(string(testTotal))
}
