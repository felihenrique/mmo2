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
	"fmt"
)

const (
	TypeNone = int16(iota)
	{{ range .Structs }}Type{{ .Name }}
	{{ end }}
)

{{ range .Structs }}
func (str *{{ .Name }}) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, Type{{ .Name }})
	buffer = serialization.Append(buffer, eventId)
	{{ range .Fields }}buffer = serialization.Append(buffer, str.{{ .Name }})
	{{ end }}
	return buffer
}

func (str *{{ .Name }}) FromBytes(data []byte) int16 {
	var n int16 = 4
	{{ range .Fields }}{{ if .BasicType }}n += serialization.Read(data[n:], &str.{{ .Name }})
	{{ else }}str.{{ .Name }} = &{{ .WithoutPointerType }}{}
	n += serialization.Read(data[n:], str.{{ .Name }}){{ end }}
	{{ end }}
	return n
}

func (str *{{ .Name }}) Type() int16 {
	return Type{{ .Name }}
}

func (str *{{ .Name }}) String() string {
	return fmt.Sprintf("{{ .Name }}: { {{ range .Fields }}{{ .Name }}: %v, {{ end }} }", {{ range .Fields }}str.{{ .Name }},{{ end }})
}
{{ end }}
`

func main() {
	newFile := strings.Replace(os.Getenv("GOFILE"), ".go", "_serialize.go", 1)
	generate.WriteFile(tpl, newFile, generate.GetData())
}
