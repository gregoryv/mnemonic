// package mnemonic provides word generators which are easily remembered.
package mnemonic

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() string {
	alfabets := []string{firstConsonants, vowels, secondConsonants}
	if startWithVowel() {
		alfabets = []string{vowels, secondConsonants, vowels}
	}
	prefix := newPrefix(alfabets...)
	return fmt.Sprintf("%s%v%v", prefix, rand.Intn(9), rand.Intn(9))
}

func startWithVowel() bool {
	return rand.Intn(2) == 1
}

func newPrefix(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

const (
	DefaultWordLength = 5

	vowels           = "aeioy"
	firstConsonants  = "bcdfghjklmnpqrstvz"
	secondConsonants = "bcdfghjklmnpqrstvxz"
)
