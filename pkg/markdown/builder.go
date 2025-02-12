package markdown

import (
	"github.com/meteormin/friday.go/pkg/markdown/component"
	"strings"
)

type Builder struct {
	components []component.Component
}

func (b *Builder) Append(c component.Component) {
	b.components = append(b.components, c)
}

func (b *Builder) NewLine() {
	txt := component.NewText()
	txt.WriteString("\n")

	b.components = append(b.components, txt)
}

func (b *Builder) HorizontalRule() {
	txt := component.NewText()
	txt.WriteString("\n---\n")

	b.components = append(b.components, txt)
}

func (b *Builder) WriteString(s string) {
	txt := component.NewText()
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
	b.components = []component.Component{}
}

func NewBuilder() *Builder {
	return &Builder{
		components: []component.Component{},
	}
}
