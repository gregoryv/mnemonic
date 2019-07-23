// package mnemonic provides word generators which are easily remembered.
package mnemonic

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() string {
	alfabets := order[:PrefixLength]
	if startWithVowel() {
		alfabets = order[1 : PrefixLength+1]
	}
	alf := make([]string, 0)
	for _, a := range alfabets {
		alf = append(alf, a)
	}
	for i := 0; i < DigitLength; i++ {
		alf = append(alf, digits)
	}
	return NewWord(alf...)
}

func startWithVowel() bool {
	return rand.Intn(2) == 1
}

func NewWord(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

const (
	PrefixLength = 3
	DigitLength  = 2

	digits           = "0123456789"
	vowels           = "aeioy"
	firstConsonants  = "bcdfghjklmnpqrstvz"
	secondConsonants = "bcdfghjklmnpqrstvxz"
)

var (
	order = []string{
		firstConsonants,
		vowels,
		secondConsonants,
		vowels,
		secondConsonants,
		vowels,
		secondConsonants,
	}
)
