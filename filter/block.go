package filter

import (
	"strings"
	"text/template"
)

type visibility string

const (
	Hide visibility = "Hide"
	Show visibility = "Show"
)

type influence string

const (
	InfluenceNone influence = "None"
)

type color string

const (
	ColorZero color = "0 0 0 0"

	ColorChrome color = "235 235 235 255"
	ColorChance color = "33 103 33 255"

	ColorBG color = "51 51 51 255" // #333
)

type rarity string

const (
	RarityNormal rarity = "Normal"
	RarityMagic  rarity = "Magic"
	RarityRare   rarity = "Rare"
	RarityUnique rarity = "Unique"
)

type cmp struct {
	OP    string
	Value string
}

func (c cmp) String() string {
	return c.OP + " " + c.Value
}

func cmpEQ(x string) *cmp {
	v := cmp{
		OP:    "==",
		Value: x,
	}
	return &v
}

func cmpLT(x string) *cmp {
	v := cmp{
		OP:    "<",
		Value: x,
	}
	return &v
}

func cmpGT(x string) *cmp {
	v := cmp{
		OP:    ">",
		Value: x,
	}
	return &v
}

type class string

func (c class) String() string {
	return `"` + string(c) + `"`
}

type Classes []string

func (c Classes) String() string {
	cs := make([]string, len(c))
	for i := range c {
		cs[i] = `"` + c[i] + `"`
	}
	return strings.Join(cs, " ")
}

type BaseTypes []string

func (bt BaseTypes) String() string {
	cs := make([]string, len(bt))
	for i := range bt {
		cs[i] = `"` + bt[i] + `"`
	}
	return strings.Join(cs, " ")
}

type block struct {
	Visibility visibility
	Continue   bool

	Class     Classes
	BaseTypes BaseTypes

	Rarity *cmp

	SocketGroup string

	Height int
	Width  int

	ElderItem    *bool
	ShaperItem   *bool
	HasInfluence influence

	BaseArmour       *cmp
	BaseEvasion      *cmp
	BaseEnergyShield *cmp

	FontSize int

	TextColor       color
	BorderColor     color
	BackgroundColor color

	DisableDropSound bool
}

func (b block) String() string {
	var buf strings.Builder
	blockTemplate.Execute(&buf, b)
	return buf.String()
}

var blockTemplate template.Template

func initBlockTemplate() {
	blockTemplate = *template.Must(template.New("section").Parse(`{{ .Visibility }}
	{{- if .Class }}
	Class {{ .Class }}{{- end}}
	{{- if .BaseTypes }}
	BaseTypes {{ .BaseTypes }}{{- end}}
	{{- if .Rarity }}
	Rarity {{ .Rarity }}{{- end }}
	{{- if .ElderItem }}
	ElderItem {{ .ElderItem }}{{- end }}
	{{- if .ShaperItem }}
	ShaperItem {{ .ShaperItem }}{{- end }}
	{{- if .HasInfluence }}
	HasInfluence "{{ .HasInfluence }}"{{- end}}
	{{- if .BaseArmour }}
	BaseArmour {{ .BaseArmour }}{{- end }}
	{{- if .BaseEvasion }}
	BaseEvasion {{ .BaseEvasion }}{{- end }}
	{{- if .BaseEnergyShield }}
	BaseEnergyShield {{ .BaseEnergyShield }}{{- end }}
	{{- if .SocketGroup }}
	SocketGroup "{{ .SocketGroup }}"{{- end }}
	{{- if .Height }}
	Height <= {{ .Height }}{{- end }}
	{{- if .Width }}
	Width <= {{ .Width }}{{- end }}
	{{- if .FontSize }}
	SetFontSize {{ .FontSize }}{{- end }}
	{{- if .TextColor }}
	SetTextColor {{ .TextColor }}{{- end }}
	{{- if .BorderColor }}
	SetBorderColor {{ .BorderColor }}{{- end }}
	{{- if .BackgroundColor }}
	SetBackgroundColor {{ .BackgroundColor }}{{- end }}
	{{- if .DisableDropSound }}
	DisableDropSound true{{- end }}
	{{- if .Continue }}
	Continue{{- end }}
`))
}
