package components

import "strings"

type TextStyle string

const (
	Bold          TextStyle = "**"
	Italic        TextStyle = "_"
	StrikeThrough TextStyle = "~~"
	BoldAndItalic TextStyle = "***"
	Subscript     TextStyle = "<sub>"
	Superscript   TextStyle = "<sup>"
	Underline     TextStyle = "<ins>"
)

type Component interface {
	String() string

	Clear()
}

type Text struct {
	sb strings.Builder
}

func (t *Text) String() string {
	if strings.HasSuffix(t.sb.String(), "\n") {
		return t.sb.String()
	} else {
		return t.sb.String() + "\n"
	}
}

func (t *Text) WriteString(s string) {
	t.sb.WriteString(s)
}

func (t *Text) InsertString(s string) {
	var newSb strings.Builder

	str := t.sb.String()
	newSb.WriteString(s)
	newSb.WriteString(str)

	t.sb = newSb
}

func (t *Text) Clear() {
	t.sb = strings.Builder{}
}

func NewText() *Text {
	return &Text{
		sb: strings.Builder{},
	}
}
