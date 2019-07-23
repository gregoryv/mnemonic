package mnemonic

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestNew(t *testing.T) {
	word := New()
	exp := PrefixLength + DigitLength
	if len(word) != exp {
		t.Error(len(word), "expected", exp)
	}
}

func Test_prefixAlfabets(t *testing.T) {
	cases := []struct {
		size int
		exp  int
	}{
		{3, 3},
		{PrefixLength, PrefixLength},
		{100, 6},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := len(prefixAlfabets(c.size))
		if got != c.exp {
			assert().Equals(got, c.exp)
		}
	}
}
