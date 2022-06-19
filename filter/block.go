package filter

import (
	"fmt"
	"strings"
	"text/template"
)

// Comparison represents equality in a criterion, can be <, <=, ==, !, >= and >.
type Comparison interface {
	renderComparison() string
}

type operator string

const (
	LT  operator = "<"
	LTE operator = "<="
	EQ  operator = "=="
	NEQ operator = "!"
	GTE operator = ">="
	GT  operator = ">"
)

type cmp struct {
	OP    operator
	Value comparable
}

func (c cmp) String() string {
	return c.renderComparison()
}

func (c cmp) renderComparison() string {
	return fmt.Sprintf("%s %v", c.OP, c.Value)
}

type comparable interface {
	isComparable()
}

func (rarity) isComparable() {}

// I is for numeric constants in comparisons.
type I uint

func (I) isComparable() {}

// Classes is a list of item classes. See https://www.poewiki.net/wiki/Item_class.
type Classes []string

func (cs Classes) String() string {
	return renderStrings([]string(cs))
}

func (cs Classes) And(xs Classes) Classes {
	return append(cs, xs...)
}

// BaseTypes is a list of item base types.
type BaseTypes []string

func (bt BaseTypes) String() string {
	return renderStrings([]string(bt))
}

func (bt BaseTypes) And(xs BaseTypes) BaseTypes {
	return append(bt, xs...)
}

func renderStrings(xs []string) string {
	cs := make([]string, len(xs))
	for i := range xs {
		cs[i] = `"` + xs[i] + `"`
	}
	return strings.Join(cs, " ")
}

type block struct {
	Visibility visibility

	Identified *bool
	Corrupted  *bool

	Class     Classes
	BaseTypes BaseTypes

	Rarity Comparison

	Quality Comparison

	AlternateQuality *bool

	SocketGroup   string
	Sockets       Comparison
	LinkedSockets Comparison

	ItemLevel Comparison
	DropLevel Comparison
	AreaLevel Comparison

	Height Comparison
	Width  Comparison

	Elder     *bool
	Shaper    *bool
	Influence influence

	Fractured   *bool
	Synthesised *bool

	GemLevel Comparison

	StackSize Comparison

	BaseArmour       Comparison
	BaseEvasion      Comparison
	BaseEnergyShield Comparison

	HasExplicitMod string

	FontSize int

	TextColor       color
	BorderColor     color
	BackgroundColor color

	DisableDropSound bool

	Continue bool
}

func (b block) String() string {
	var buf strings.Builder
	blockTemplate.Execute(&buf, b)
	return buf.String()
}

var blockTemplate template.Template

func init() {
	blockTemplate = *template.Must(template.New("section").Parse(`{{ .Visibility }}
	{{- if .Identified }}
	Identified {{ .Identified }}{{- end }}
	{{- if .Corrupted }}
	Corrupted {{ .Corrupted }}{{- end }}
	{{- if .Class }}
	Class {{ .Class }}{{- end }}
	{{- if .BaseTypes }}
	BaseType {{ .BaseTypes }}{{- end }}
	{{- if .Rarity }}
	Rarity {{ .Rarity }}{{- end }}
	{{- if .Quality }}
	Quality {{ .Quality }}{{- end }}
	{{- if .AlternateQuality }}
	AlternateQuality {{ .AlternateQuality }}{{- end }}
	{{- if .ItemLevel }}
	ItemLevel {{ .ItemLevel }}{{- end }}
	{{- if .DropLevel }}
	DropLevel {{ .DropLevel }}{{- end }}
	{{- if .AreaLevel }}
	AreaLevel {{ .AreaLevel }}{{- end }}
	{{- if .GemLevel }}
	GemLevel {{ .GemLevel }}{{- end }}

	{{- if .Elder }}
	ElderItem {{ .Elder }}{{- end }}
	{{- if .Shaper }}
	ShaperItem {{ .Shaper }}{{- end }}
	{{- if .Influence }}
	HasInfluence "{{ .Influence }}"{{- end}}
	{{- if .Fractured }}
	FracturedItem {{ .Fractured }}{{- end }}
	{{- if .Synthesised }}
	SynthesisedItem {{ .Synthesised }}{{- end }}
	{{- if .BaseArmour }}
	BaseArmour {{ .BaseArmour }}{{- end }}
	{{- if .BaseEvasion }}
	BaseEvasion {{ .BaseEvasion }}{{- end }}
	{{- if .BaseEnergyShield }}
	BaseEnergyShield {{ .BaseEnergyShield }}{{- end }}
	{{- if .SocketGroup }}
	SocketGroup "{{ .SocketGroup }}"{{- end }}
	{{- if .Sockets }}
	Sockets {{ .Sockets }}{{- end }}
	{{- if .LinkedSockets }}
	LinkedSockets {{ .LinkedSockets }}{{- end }}
	{{- if .Height }}
	Height {{ .Height }}{{- end }}
	{{- if .Width }}
	Width {{ .Width }}{{- end }}
	{{- if .HasExplicitMod }}
	HasExplicitMod {{ .HasExplicitMod }}{{- end }}
	{{- if .StackSize }}
	StackSize {{ .StackSize }}{{- end }}
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
