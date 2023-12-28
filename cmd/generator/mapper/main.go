package main

import (
	"mmo2/pkg/generate"
	"os"
	"strings"
)

var tpl = `
package {{ .Package }}

import (
	"mmo2/pkg/serialization"
)

type unmarshal = func(data []byte) (serialization.ISerializable, int16)

var Mapper = []unmarshal{
{{ range .Structs }}func(data []byte) (serialization.ISerializable, int16) {
		return Parse{{ .Name }}(data)
	},
{{ end }}
}
`

func main() {
	newFile := strings.Replace(os.Getenv("GOFILE"), ".go", "_mapper.go", 1)
	generate.WriteFile(tpl, newFile, generate.GetData())
}
