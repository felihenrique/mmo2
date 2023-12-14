package main

import (
	"regexp"
	"strings"
)

type Field struct {
	Name string
	Type string
}

type Struct struct {
	Name string
}

func IsSlice(typeStr string) bool {
	return strings.HasPrefix("[]", typeStr)
}

func Tokenize(file string) []string {
	re := regexp.MustCompile(`\n|\t`)
	file = re.ReplaceAllString(file, " ")

	tokens := strings.Fields(file)

	var nonEmptyTokens []string
	for _, token := range tokens {
		if token != "" {
			nonEmptyTokens = append(nonEmptyTokens, token)
		}
	}

	return nonEmptyTokens
}

// PACKAGE, READED TOKENS
func ReadPackage(tokens []string) (string, int) {
	return tokens[1], 2
}
