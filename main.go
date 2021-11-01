package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
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

		// Initialize write to print columns
		writer := tabwriter.NewWriter(os.Stdout, 3, 3, 1, ' ', tabwriter.AlignRight)

		// Print the header (inidividual letters of the word)
		wordDisplay := getWordDisplay(&word)
		fmt.Fprintln(writer, wordDisplay)

		// Print a seperator
		seperator := strings.Repeat("---\t", len(word.originalString))
		fmt.Fprintln(writer, seperator)

		// Print a row for each cipher
		for _, cipher := range ciphers {
			fmt.Fprintln(writer, getWordValuesDisplay(&word, cipher))
		}

		writer.Flush()
	}
}
