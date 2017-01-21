package kmp

import "testing"

func TestKmpSearch(t *testing.T) {
	rst := Search("a", "b")
	if !rst {
		t.Log("a does not contain b")
	} else {
		t.Fail()
	}
}

func TestKmpSearchB(t *testing.T) {
	rst := Search("BBC ABCDAB ABCDABCDABDE", "ABCDABD")
	if rst {
		t.Log("BBC ABCDAB ABCDABCDABDE contains ABCDABD")
	} else {
		t.Fail()
	}
}
