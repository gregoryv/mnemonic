package main

import (
	"fmt"
	"time"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/mnemonic"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		digits = cli.Option("-d, --digits",
			"number of digits after prefix").Int(2)

		prefixLen = cli.Option("-p, --prefix-len",
			"number of letters starting the word").Int(3)

		repeat = cli.Option("-r, --repeat",
			"repeat generating words until interrupted with ctrl+c").Bool()
	)
	cli.Parse()

	fmt.Println(mnemonic.New(prefixLen, digits))
	if repeat {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println(mnemonic.New(prefixLen, digits))
		}
	}
}
