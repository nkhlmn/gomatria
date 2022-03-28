package main

import (
	"fmt"
	"unicode"
	"regexp"
	"errors"
	"strings"
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

func getAudioSigil(word *Word, cipherType Cipher) string {
	var notes []string
	for _, char := range word.originalString {
		letterValue := getLetterValue(&char, cipherType)
		letterNote := numberToNote(&letterValue)
		notes = append(notes, letterNote)
	}
	return strings.Join(notes[:], " ")
}

func numberToNote(num *int) string {
	var notes [12]string
	useSharps := false
	if useSharps {
		notes = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	} else {
		notes = [12]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}

	octaveOffset := 1
	octave := (*num / 12) + octaveOffset

	var noteNumber int
	if *num == 12 {
		noteNumber = 12
	} else {
		noteNumber = *num % 12
	}
	note := notes[noteNumber - 1]

	return fmt.Sprintf("%s%v", note, octave)
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
		result += fmt.Sprintf("%v\t", string(char))
	}
	return result
}

func getWordValuesDisplay(w *Word, cipherType Cipher) string {
	var result string
	for _, char := range w.originalString {
		letterValue := getLetterValue(&char, cipherType)
		result += fmt.Sprintf("%v\t", letterValue)
	}
	result += fmt.Sprintf(
		"\t(%v) %v",
		w.CipherValue(cipherType),
		cipherType.String(),
	)
	return result
}
