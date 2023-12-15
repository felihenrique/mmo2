package main

import (
	"mmo2/pkg/generate"
	"os"
	"strings"
)

var tpl = `
package {{ .Package }}

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	{{ range .Structs }}Type{{ .Name }}
	{{ end }}
)

{{ range .Structs }}
func (str *{{ .Name }}) ToBytes() []byte {
	buffer := make([]byte, 0)
	{{ range .Fields }}buffer = serialization.WriteBinary(buffer, str.{{ .Name }})
	{{ end }}
	return buffer
}

func (str *{{ .Name }}) FromBytes(data []byte) int16 {
	var n int16 = 0
	{{ range .Fields }}n += serialization.ReadBinary(data[n:], &str.{{ .Name }})
	{{ end }}
	return n
}

func (str *{{ .Name }}) EvType() int16 {
	return Type{{ .Name }}
}
{{ end }}
`

func main() {
	newFile := strings.Replace(os.Getenv("GOFILE"), ".go", "_serialize.go", 1)
	generate.WriteFile(tpl, newFile, generate.GetData())
}