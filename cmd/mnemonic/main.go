package main

import (
	"flag"

	"github.com/gregoryv/stamp"
	"xwing.7de.se/mnemonic"
)

//go:generate stamp -go build_stamp.go -clfile ../../CHANGELOG.md
func main() {
	me := &mnemonic.MainEntryPoint{}
	me.InitFlags()
	stamp.InitFlags()
	flag.Parse()
	stamp.AsFlagged()
	me.Enter()
	me.Exit()
}
