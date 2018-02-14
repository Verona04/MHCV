package main

import "testing"

func TestExtendedASCIIText(t * testing.T) {
	var testStr = "hæ"
	testResult := ExtendedASCIIText(testStr)

	if len(testResult) == 0 {
		t.Fail()
	}
}