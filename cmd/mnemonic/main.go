package main

import (
	"flag"
	"time"

	"github.com/gregoryv/stamp"
	"xwing.7de.se/mnemonic"
)

//go:generate stamp -go build_stamp.go -clfile ../../CHANGELOG.md
func main() {
	// Add -v and -vv flags
	stamp.InitFlags()

	flag.IntVar(
		&mnemonic.PrefixLength, "p",
		mnemonic.PrefixLength, "number of letters starting the word",
	)
	flag.IntVar(
		&mnemonic.DigitLength, "d",
		mnemonic.DigitLength, "number of digits after prefix",
	)
	var repeat bool
	flag.BoolVar(
		&repeat, "r",
		repeat, "repeat generating words until interrupted with ctrl+c",
	)
	flag.Parse()
	stamp.AsFlagged()

	print(mnemonic.New())
	if repeat {
		for {
			time.Sleep(300 * time.Millisecond)
			print("\n", mnemonic.New())
		}
	}
}
