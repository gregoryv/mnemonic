package mnemonic

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestMainEntryPoint(t *testing.T) {
	cases := []struct {
		me  *MainEntryPoint
		exp ExitCode
	}{
		{
			&MainEntryPoint{},
			ExitOk,
		},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		c.me.Enter()
		got := c.me.exitCode
		assert().Equals(got, c.exp)
	}
}
