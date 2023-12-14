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
	Name   string
	Fields []Field
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
func ReadPackage(tokens []string) string {
	return tokens[1]
}

func ReadFields(tokens []string) ([]Field, int) {
	fields := make([]Field, 0)
	for i := 0; ; i += 2 {
		if tokens[i] == "}" {
			break
		}
		fields = append(fields, Field{
			Name: tokens[i],
			Type: tokens[i+1],
		})
	}
	return fields, len(fields) * 2
}

func ReadNextStruct(tokens []string) (Struct, int) {
	pos := 0
	str := Struct{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "type" || tokens[i+1] == "Ack" {
			continue
		}
		str.Name = tokens[i+1]
		fields, readed := ReadFields(tokens[i+4:])
		str.Fields = fields
		pos = i + readed + 4
		break
	}
	return str, pos
}
