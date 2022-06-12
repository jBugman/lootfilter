package filter

import (
	"strings"
	"text/template"
)

const defaltFontSize = 34

type Filter struct {
	sections []section

	Chromatic bool
}

func (f Filter) virtualSections() []section {
	ss := f.sections
	if f.Chromatic {
		ss = append(ss,
			section{
				Visibility:  Show,
				SocketGroup: "RGB",
				Height:      3,
				Width:       1,
				BorderColor: "235 235 235 255",
			},
			section{
				Visibility:  Show,
				SocketGroup: "RGB",
				Height:      2,
				Width:       2,
				BorderColor: "235 235 235 255",
			},
		)
	}
	return ss
}

func (f Filter) String() string {
	var buf strings.Builder
	for _, sec := range f.virtualSections() {
		_, _ = buf.WriteString(sec.String())
		_, _ = buf.WriteString("\n")
	}
	res := strings.TrimSpace(buf.String())
	res += "\n"
	return res
}

var defaultFilter = Filter{
	sections: []section{
		{
			Visibility: Show,
			FontSize:   34,
		},
	},
	Chromatic: true,
}

var Default = defaultFilter

type visibility string

const Hide visibility = "Hide"
const Show visibility = "Show"

type section struct {
	Visibility visibility

	SocketGroup string

	Height int
	Width  int

	FontSize    int
	BorderColor string
}

func (sec section) String() string {
	var buf strings.Builder
	sectionTemplate.Execute(&buf, sec)
	return buf.String()
}

var sectionTemplate template.Template

func init() {
	sectionTemplate = *template.Must(template.New("section").Parse(`{{ .Visibility }}
	{{- if .SocketGroup }}
	SocketGroup "{{ .SocketGroup }}"{{- end }}
	{{- if .Height }}
	Height <= {{ .Height }}{{- end }}
	{{- if .Width }}
	Width <= {{ .Width }}{{- end }}
	{{- if .FontSize }}
	SetFontSize {{ .FontSize }}{{- end }}
	{{- if .BorderColor }}
	SetBorderColor {{ .BorderColor }}{{- end }}
`))
}
