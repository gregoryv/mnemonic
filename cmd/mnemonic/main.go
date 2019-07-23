package main

import (
	"flag"

	"xwing.7de.se/mnemonic"
)

func main() {
	flag.IntVar(
		&mnemonic.PrefixLength, "p",
		mnemonic.PrefixLength, "number of letters starting the word",
	)
	flag.IntVar(
		&mnemonic.DigitLength, "d",
		mnemonic.DigitLength, "number of digits after prefix",
	)
	flag.Parse()
	print(mnemonic.New())
}
