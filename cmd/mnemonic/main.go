package main

import (
	"flag"
	"os"

	"xwing.7de.se/mnemonic"
)

func main() {
	me := &mnemonic.MainEntryPoint{}
	me.InitFlags()
	flag.Usage = me.Usage(os.Stderr)
	flag.Parse()
	me.Enter()
	me.Exit()
}
