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
	alfabets := prefixAlfabets(PrefixLength)
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

func prefixAlfabets(size int) []string {
	if size >= len(order) {
		size = len(order) - 1
	}
	alfabets := order[:size]
	if startWithVowel() {
		alfabets = order[1 : size+1]
	}
	return alfabets
}

func NewWord(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

var (
	PrefixLength = 3
	DigitLength  = 2
)

const (
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
