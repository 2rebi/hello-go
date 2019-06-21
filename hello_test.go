package hello

import "testing"

func TestGet(t *testing.T) {
	word := "World"
	h := Hello{Word: word}
	if h.get() != ("Hello " + word) {
		t.Error("Test Failed")
	}
}
