// Package mnemonic provides word generators for easily remembered and
// pronouncible.
package mnemonic

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New returns a mnemonic
func New(prefixLen, digitLen int) string {
	alfabets := prefixAlfabets(prefixLen)
	alf := make([]string, 0)
	for _, a := range alfabets {
		alf = append(alf, a)
	}
	for i := 0; i < digitLen; i++ {
		alf = append(alf, Digits)
	}
	return NewWord(alf...)
}

// prefixAlfabets returns a set of letter character strings in a
// specific order to be used while generating a word.
func prefixAlfabets(size int) []string {
	if size >= len(Order) {
		size = len(Order) - 1
	}
	alfabets := Order[:size]
	startWithVowel := rand.Intn(2) == 1
	if startWithVowel {
		alfabets = Order[1 : size+1]
	}
	return alfabets
}

// NewWord returns a random string based on one letter from each alfabet.
func NewWord(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

// todo make these public and reusable, to e.g. also combine uppercase letters
const (
	Digits           = "0123456789"
	Vowels           = "aeioy"
	FirstConsonants  = "bcdfghjklmnpqrstvz"
	SecondConsonants = "bcdfghjklmnpqrstvxz"
)

var (
	// todo, generate instead in func New based on len
	Order = []string{
		FirstConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
		Vowels,
		SecondConsonants,
	}
)
