package filter

import (
	"strings"
	"text/template"
)

const minFontSize = 18
const clickableFontSize = 26

type visibility string

const (
	// Hide means block is about hiding loot.
	Hide visibility = "Hide"
	// Show means block is about displaying loot.
	Show visibility = "Show"
)

type influence string

const (
	// InfluenceNone is for no-influence items.
	InfluenceNone influence = "None"
)

type color string

const (
	// ColorZero is for invisibility.
	ColorZero color = "0 0 0 0"

	// ColorCurrency is PoE native color for currency.
	ColorCurrency color = "170 158 130 255"

	// ColorChrome is NeverSink's default color for the chromatic recipe.
	ColorChrome color = "235 235 235 255"
	// ColorChance is NeverSink's default color for chancing bases.
	ColorChance color = "33 103 33 255"
	// Color6S is  NeverSink's default color for 6S bases.
	Color6S color = "200 200 200 255"

	// ColorBG is default background color.
	ColorBG color = "51 51 51 255" // #333
)

type rarity string

const (
	// RarityNormal is for white items.
	RarityNormal rarity = "Normal"
	// RarityMagic is for magic items.
	RarityMagic rarity = "Magic"
	// RarityRare is for rare items.
	RarityRare rarity = "Rare"
	// RarityUnique is for unique items.
	RarityUnique rarity = "Unique"
)

// Comparison represents equality in a criterion, can be <, <=, ==, >= and >.
type Comparison interface {
	RenderComparison() string
}

type cmp struct {
	OP    string
	Value string
}

func (c cmp) String() string {
	return c.RenderComparison()
}

func (c *cmp) RenderComparison() string {
	return c.OP + " " + c.Value
}

// CmpLT is a Comparison for <.
func CmpLT(x string) Comparison {
	v := cmp{
		OP:    "<",
		Value: x,
	}
	return &v
}

// CmpLTE is a Comparison for <=.
func CmpLTE(x string) Comparison {
	v := cmp{
		OP:    "<=",
		Value: x,
	}
	return &v
}

// CmpEQ is a Comparison for ==.
func CmpEQ(x string) Comparison {
	v := cmp{
		OP:    "==",
		Value: x,
	}
	return &v
}

// CmpNEQ is a Comparison for !.
func CmpNEQ(x string) Comparison {
	v := cmp{
		OP:    "!",
		Value: x,
	}
	return &v
}

// CmpGTE is a Comparison for >=.
func CmpGTE(x string) Comparison {
	v := cmp{
		OP:    ">=",
		Value: x,
	}
	return &v
}

// CmpGT is a Comparison for >.
func CmpGT(x string) Comparison {
	v := cmp{
		OP:    ">",
		Value: x,
	}
	return &v
}

// Classes is a list of item classes. See https://www.poewiki.net/wiki/Item_class.
type Classes []string

func (c Classes) String() string {
	cs := make([]string, len(c))
	for i := range c {
		cs[i] = `"` + c[i] + `"`
	}
	return strings.Join(cs, " ")
}

// BaseTypes is a list of item base types.
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

	Rarity Comparison

	Quality Comparison

	SocketGroup   string
	Sockets       Comparison
	LinkedSockets Comparison

	ItemLevel Comparison
	DropLevel Comparison
	AreaLevel Comparison

	Height int
	Width  int

	ElderItem    *bool
	ShaperItem   *bool
	HasInfluence influence

	BaseArmour       Comparison
	BaseEvasion      Comparison
	BaseEnergyShield Comparison

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
	BaseType {{ .BaseTypes }}{{- end}}
	{{- if .Rarity }}
	Rarity {{ .Rarity }}{{- end }}
	{{- if .Quality }}
	Quality {{ .Quality }}{{- end }}
	{{- if .ItemLevel }}
	ItemLevel {{ .ItemLevel }}{{- end }}
	{{- if .DropLevel }}
	DropLevel {{ .DropLevel }}{{- end }}
	{{- if .AreaLevel }}
	AreaLevel {{ .AreaLevel }}{{- end }}
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
	{{- if .Sockets }}
	Sockets {{ .Sockets }}{{- end }}
	{{- if .LinkedSockets }}
	LinkedSockets {{ .LinkedSockets }}{{- end }}
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
