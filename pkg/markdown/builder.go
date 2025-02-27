package markdown

import (
	"github.com/meteormin/friday.go/pkg/markdown/components"
	"strings"
)

type Builder struct {
	components []components.Component
}

func (b *Builder) Append(c components.Component) {
	b.components = append(b.components, c)
}

func (b *Builder) NewLine() {
	txt := components.NewText()
	txt.WriteString("\n")

	b.components = append(b.components, txt)
}

func (b *Builder) HorizontalRule() {
	txt := components.NewText()
	txt.WriteString("\n---\n")

	b.components = append(b.components, txt)
}

func (b *Builder) WriteString(s string) {
	txt := components.NewText()
	txt.WriteString(s)

	b.components = append(b.components, txt)
}

func (b *Builder) String() string {
	var sb strings.Builder
	for _, c := range b.components {
		sb.WriteString(c.String())
	}
	return sb.String()
}

func (b *Builder) Clear() {
	b.components = []components.Component{}
}

func NewBuilder() *Builder {
	return &Builder{
		components: []components.Component{},
	}
}
