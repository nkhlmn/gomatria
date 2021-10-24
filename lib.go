package main

import (
	"fmt"
	"unicode"
	"regexp"
	"errors"
)

func validateWordString(wordString *string) (bool, error) {
	var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(*wordString)
	if (isAlphaNumeric) {
		return true, nil
	} else {
		return false, errors.New("Args must only contain letters and numbers!")
	}
}

func getWord(wordString string) Word {
	var result = Word{
		originalString: wordString,
	}
	return result
}

func getLetterValue(c *rune, cipherType Cipher) int {
	if (unicode.IsDigit(*c)) {
		return int(*c - '0')
	}
	char := unicode.ToLower(*c) - 96
	var result int
	switch cipherType {
	case Ordinal:
		result = int(char)
	case OrdinalReverse:
		result = 27 - int(char)
	case FullReduction:
		if int(char)%9 == 0 {
			result += 9
		} else {
			result = int(char) % 9
		}
	case FullReductionReverse:
		result = 27 - int(char)
		if result%9 == 0 {
			result = 9
		} else {
			result = result % 9
		}
	}
	return result
}

func getWordValue(word *Word, cipherType Cipher) int {
	var value int
	for _, char := range word.originalString {
		letterValue := getLetterValue(&char, cipherType)
		value += int(letterValue)
	}
	return value
}

func getWordDisplay(w *Word) string {
	var result string
	for _, char := range w.originalString {
		result += fmt.Sprintf("%-5v", string(char))
	}
	return result
}

func getWordValuesDisplay(w *Word, cipherType Cipher) string {
	var result string
	for _, char := range w.originalString {
		letterValue := getLetterValue(&char, cipherType)
		result += fmt.Sprintf("%-5v", letterValue)
	}
	result += fmt.Sprintf(
		"(%v) %v",
		getWordValue(w, cipherType),
		cipherType.String(),
	)
	return result
}
