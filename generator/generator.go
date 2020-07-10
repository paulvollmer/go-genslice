package generator

import (
	"bytes"
	"sort"
	"text/template"
)

var DefaultTemplate = `// Generated code. Do not edit!

package {{.PackageName}}

var {{.VarName}} = []{{.VarType}}{
{{range $i, $val := .Data}}	"{{$val}}", // {{$i}}
{{end}}
}
`

type Config struct {
	Data        []string
	PackageName string
	VarType     string
	VarName     string
	Sort        bool
}

func GenerateCode(tpl string, config Config) ([]byte, error) {
	t, err := template.New("gen").Parse(tpl)
	if err != nil {
		return nil, err
	}

	if config.Sort {
		sort.Strings(config.Data)
	}

	var w bytes.Buffer
	err = t.Execute(&w, config)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
