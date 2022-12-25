package main

import "testing"

func TestMultiple(t *testing.T) {
	var x, y, result = 2, 2, 4

	testResult := multiple(x, y)

	if testResult != result {
		t.Errorf("%d != %d", result, testResult)
	}
}
