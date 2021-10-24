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
	var arg = os.Args[1]
	var isValid, err = validateWordString(&arg)
	if (err != nil) {
		fmt.Println(err)
	} else if (isValid) {
		var word = getWord(arg)
		var ciphers = []Cipher{
			Ordinal,
			FullReduction,
			OrdinalReverse,
			FullReductionReverse,
		}

		wordDisplay := getWordDisplay(&word)
		fmt.Println(wordDisplay)
		seperator := strings.Repeat("-", len(wordDisplay))
		fmt.Println(seperator)
		for _, cipher := range ciphers {
			fmt.Println(getWordValuesDisplay(&word, cipher))
		}
	}
}
