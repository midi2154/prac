package main

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {

	words := strings.Fields(text)

	words = ApplyHex(words)
	words = ApplyBin(words)
	words = ApplyCase(words)

	result := strings.Join(words, " ")

	result = FixPunctuation(result)
	result = FixQuotes(result)
	result = FixArticles(result)

	return result
}

func ApplyHex(words []string) []string {

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {

			value, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
			}

			words = RemoveWord(words, i)
			i--
		}
	}
	return words
}

func ApplyBin(words []string) []string {

	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {

			value, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
			}

			words = RemoveWord(words, i)
			i--
		}
	}
	return words
}

func ApplyCase(words []string) []string {

	for i := 0; i < len(words); i++ {

		// Check if it contains comma version like (up, 3)
		if strings.HasPrefix(words[i], "(up") ||
			strings.HasPrefix(words[i], "(low") ||
			strings.HasPrefix(words[i], "(cap") {

			mode, count := ParseCaseInstruction(words[i])

			if i > 0 {

				for j := 0; j < count && i-1-j >= 0; j++ {

					index := i - 1 - j

					if mode == "up" {
						words[index] = strings.ToUpper(words[index])
					}
					if mode == "low" {
						words[index] = strings.ToLower(words[index])
					}
					if mode == "cap" {
						words[index] = Capitalize(words[index])
					}
				}

				words = RemoveWord(words, i)
				i--
			}
		}
	}

	return words
}

func ParseCaseInstruction(word string) (string, int) {

	word = strings.Trim(word, "()")

	parts := strings.Split(word, ",")

	mode := parts[0]
	count := 1

	if len(parts) > 1 {
		number := strings.TrimSpace(parts[1])
		n, err := strconv.Atoi(number)
		if err == nil {
			count = n
		}
	}

	return mode, count
}

func Capitalize(word string) string {

	if len(word) == 0 {
		return word
	}

	first := strings.ToUpper(string(word[0]))
	rest := strings.ToLower(word[1:])

	return first + rest
}

func RemoveWord(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func FixPunctuation(text string) string {

	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {

		// Remove space before punctuation
		text = strings.ReplaceAll(text, " "+p, p)

		// Ensure one space after punctuation
		text = strings.ReplaceAll(text, p+" ", p+" ")
	}

	return text
}

func FixQuotes(text string) string {

	for strings.Contains(text, "' ") {
		text = strings.ReplaceAll(text, "' ", "'")
	}

	for strings.Contains(text, " '") {
		text = strings.ReplaceAll(text, " '", "'")
	}

	return text
}

func FixArticles(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		current := words[i]
		next := words[i+1]

		if current == "a" || current == "A" {

			first := strings.ToLower(string(next[0]))

			if first == "a" ||
				first == "e" ||
				first == "i" ||
				first == "o" ||
				first == "u" ||
				first == "h" {

				if current == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
