package main

import (
	"flag"
	"os"

	"github.com/gregoryv/stamp"
	"xwing.7de.se/mnemonic"
)

//go:generate stamp -go build_stamp.go -clfile ../../changelog.md
func main() {
	me := &mnemonic.MainEntryPoint{}
	me.InitFlags()
	stamp.InitFlags()
	flag.Usage = me.Usage(os.Stderr)
	flag.Parse()
	stamp.AsFlagged()
	me.Enter()
	me.Exit()
}
