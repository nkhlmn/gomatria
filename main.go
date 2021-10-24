package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var argsLength = len(os.Args[1:])
	if argsLength < 1 {
		fmt.Println("Must provide an argument!")
		os.Exit(1)
	}
	var wordArg = os.Args[1]
	var gematria = getWord(wordArg)
	var ciphers = []Cipher{
		Ordinal,
		FullReduction,
		OrdinalReverse,
		FullReductionReverse,
	}

	wordDisplay := getWordDisplay(&gematria)
	fmt.Println(wordDisplay)
	seperator := strings.Repeat("-", len(wordDisplay))
	fmt.Println(seperator)
	for _, cipher := range ciphers {
		fmt.Println(getWordValuesDisplay(&gematria, cipher))
	}
}
