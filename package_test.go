package mnemonic

import (
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/qual"
)

func TestQuality(t *testing.T) {
	qual.Standard(t)
}

func TestNew(t *testing.T) {
	prefixLen := 4
	digitLen := 2
	word := New(prefixLen, digitLen)
	exp := prefixLen + digitLen
	if len(word) != exp {
		t.Error(len(word), "expected", exp)
	}
}

func Test_prefixAlfabets(t *testing.T) {
	maxPrefixLen := len(order) - 1
	cases := []struct {
		size int
		exp  int
	}{
		{3, 3},
		{100, maxPrefixLen},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := len(prefixAlfabets(c.size))
		if got != c.exp {
			assert().Equals(got, c.exp)
		}
	}
}
