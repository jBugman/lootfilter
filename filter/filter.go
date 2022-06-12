package filter

import (
	"strings"
	"text/template"
)

const defaltFontSize = 34

type Filter struct {
	sections []section
}

func (f Filter) String() string {
	var buf strings.Builder
	for _, sec := range f.sections {
		sectionTemplate.Execute(&buf, sec)
	}
	return buf.String()
}

var Default = Filter{
	sections: []section{
		{
			Visibility: show,
			FontSize:   34,
		},
	},
}

type visibility string

const hide visibility = "Hide"
const show visibility = "Show"

type section struct {
	Visibility visibility
	FontSize   int
}

var sectionTemplate template.Template

func init() {
	sectionTemplate = *template.Must(template.New("section").Parse(`{{ .Visibility }}
	{{ if .FontSize -}}SetFontSize {{ .FontSize }}{{- end }}
`))
}
