package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

type Data struct {
	Filename string
	Package  string
	Structs  []Struct
}

func formatFile(file string) ([]byte, error) {
	command := exec.Command("bash", "-c", fmt.Sprintf("gofmt -w %s", file))
	output, err := command.CombinedOutput()
	return output, err
}

func WriteFile(data Data) {
	tplSrc := `
package {{ .Package }}

const (
	TypeAck = int16(iota + 1)
	{{ range .Structs }}Type{{ .Name }}
	{{ end }}
)

{{ range .Structs }}
func (str *{{ .Name }}) toBytes() []byte {
	buffer := make([]byte, 0)
	{{ range .Fields }}buffer = WriteBinary(buffer, str.{{ .Name }})
	{{ end }}
	return buffer
}

func (str *{{ .Name }}) fromBytes(data []byte) int16 {
	var n int16 = 0
	{{ range .Fields }}n += ReadBinary(data[n:], &str.{{ .Name }})
	{{ end }}
	return n
}

func (str *{{ .Name }}) evType() int16 {
	return Type{{ .Name }}
}
{{ end }}
`
	tpl, err := template.New("render").Parse(tplSrc)
	if err != nil {
		panic(err)
	}
	newFile := strings.Replace(data.Filename, ".go", "_gen.go", 1)
	file, err := os.Create(newFile)
	tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
	_, err = formatFile(newFile)
	if err != nil {
		panic(err)
	}
}

func main() {
	file := os.Getenv("GOFILE")

	if len(file) == 0 {
		fmt.Println("Env GOFILE is required")
		return
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
		readed += n
		if readed+1 >= len(tokens) {
			break
		}
		strs = append(strs, str)
	}

	WriteFile(Data{
		Filename: file,
		Package:  ReadPackage(tokens),
		Structs:  strs,
	})
}
