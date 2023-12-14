package main

import (
	"mmo2/pkg/generate"
)

var tpl = `
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

func main() {

	generate.WriteFile(tpl, generate.GetData())
}
