package main

import (
	"fmt"
	"unicode"
)

type TextColor int

const (
	TextColorRed TextColor = iota
	TextColorGreen
	TextColorYellow
	TextColorBlue
	TextColorPink
)

func makeTextColor(text string, color TextColor) string {
	switch color {
	case TextColorRed:
		return fmt.Sprintf("\033[31m%s\033[0m", text)
	case TextColorGreen:
		return fmt.Sprintf("\033[32m%s\033[0m", text)
	case TextColorYellow:
		return fmt.Sprintf("\033[33m%s\033[0m", text)
	case TextColorBlue:
		return fmt.Sprintf("\033[34m%s\033[0m", text)
	case TextColorPink:
		return fmt.Sprintf("\033[38;5;205m%s\033[0m", text)
	default:
		return text
	}
}

func isOnlyLetters(value string) bool {
	for _, char := range value {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true;
}
