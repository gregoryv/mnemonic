package mnemonic

import "testing"

func TestNew(t *testing.T) {
	word := New()
	if len(word) != DefaultWordLength {
		t.Error(len(word), "expected", DefaultWordLength)
	}
}
