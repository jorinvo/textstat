package main

import (
	"os/exec"
	"testing"
)

var expected = `
  total words: 904.061
  unique words: 59.723
  average word length: 4
  most used words: the, and, i, to, of, a, my, in, you, that
`

func TestBinary(t *testing.T) {
	out, err := exec.Command("$GOPATH/bin/textstat-go", "texts/shakespeare.txt").Output()
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	result := string(out)
	if result != expected {
		t.Logf("Expected %s", expected)
		t.Logf("Got %s", result)
		t.Fail()
	}
}
