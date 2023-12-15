package generate

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

type Data struct {
	Filename string
	Package  string
	Structs  []Struct
}

type Field struct {
	Name string
	Type string
}

type Struct struct {
	Name   string
	Fields []Field
}

func formatFile(file string) ([]byte, error) {
	command := exec.Command("bash", "-c", fmt.Sprintf("gofmt -w %s", file))
	output, err := command.CombinedOutput()
	return output, err
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
		if tokens[i] != "type" {
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

func WriteFile(tplSrc string, newFile string, data Data) {
	tpl, err := template.New("render").Parse(tplSrc)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(newFile)
	tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
	_, err = formatFile(newFile)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func GetData() Data {
	file := os.Getenv("GOFILE")

	if len(file) == 0 {
		panic("Env GOFILE is required")
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	tokens := Tokenize(string(bytes))

	strs := make([]Struct, 0)
	readed := 0
	for {
		str, n := ReadNextStruct(tokens[readed:])
		if str.Name != "" {
			strs = append(strs, str)
		}
		readed += n
		if readed+1 >= len(tokens) {
			break
		}
	}

	return Data{
		Filename: file,
		Package:  ReadPackage(tokens),
		Structs:  strs,
	}
}
