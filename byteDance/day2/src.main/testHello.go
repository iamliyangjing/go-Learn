package main

import "testing"

func HelloTom() string {
	return "Jerry"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Error("%s do not match actual %s", expectOutput, output)
	}
}
