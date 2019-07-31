package mnemonic

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type MainEntryPoint struct {
	PrefixLen int
	DigitLen  int
	Repeat    bool

	output   io.Writer
	exitCode ExitCode
}

func (e *MainEntryPoint) Usage(w io.Writer) func() {
	return func() {
		fmt.Fprintf(w, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprint(w, "Options\n")
		flag.PrintDefaults()
	}
}

func (e *MainEntryPoint) InitFlags() {
	flag.IntVar(
		&e.PrefixLen, "p",
		3, "number of letters starting the word",
	)
	flag.IntVar(
		&e.DigitLen, "d",
		2, "number of digits after prefix",
	)

	flag.BoolVar(
		&e.Repeat, "r",
		false, "repeat generating words until interrupted with ctrl+c",
	)
}

func (e *MainEntryPoint) String() string {
	return fmt.Sprintf(
		"MainEntryPoint: p=%v, d=%v, r=%v, exitCode=%v",
		e.PrefixLen, e.DigitLen, e.Repeat, e.exitCode,
	)
}

func (e *MainEntryPoint) Enter() {
	e.setDefaults()
	fmt.Fprint(e.output, New(e.PrefixLen, e.DigitLen))
	if e.Repeat {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Fprint(e.output, "\n", New(e.PrefixLen, e.DigitLen))
		}
	}
}

func (e *MainEntryPoint) setDefaults() {
	if e.output == nil {
		e.output = os.Stdout
	}
	e.exitCode = ExitOk
}

func (e *MainEntryPoint) fail(err error) {
	fmt.Fprint(os.Stderr, err)
	e.exitCode = ExitFail
}

func (e *MainEntryPoint) Exit() {
	os.Exit(e.exitCode)
}

type ExitCode = int

const (
	ExitOk ExitCode = iota
	ExitFail
)
