package mnemonic

import "testing"

func TestNew(t *testing.T) {
	word := New()
	exp := PrefixLength + DigitLength
	if len(word) != exp {
		t.Error(len(word), "expected", exp)
	}
}
