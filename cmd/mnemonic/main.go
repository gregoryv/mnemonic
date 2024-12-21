package main

import (
	"fmt"
	"sort"
	"strings"

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

		count = cli.Option("-c, --count",
			"number of generated mnemonics").Int(1)

		begin = cli.Option("-b, --begin-with",
			"must start with letter").String("")
	)
	cli.Parse()

	// generate mnemonics
	result := make([]string, 0, count)
	for {
		word := mnemonic.New(prefixLen, digits)
		if !strings.HasPrefix(word, begin) {
			continue
		}
		result = append(result, word)
		if len(result) == count {
			break
		}
	}
	// sort them
	sort.Strings(result)

	// print them to stdout
	last := result[0][0]
	maxLineWords := 8
	var n int
	for _, word := range result {
		n++
		if v := word[0]; v != last || n == maxLineWords {
			fmt.Println()
			last = v
			n = 0
		}
		fmt.Print(word, " ")
	}
	// final new line only if needed
	if count%2 != 0 || n < count / 8{
		fmt.Println()
	}
}
