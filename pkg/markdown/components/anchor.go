package components

import (
	"fmt"
)

type Anchor struct {
	*Text
}

func NewAnchor(text string) *Anchor {
	a := &Anchor{
		Text: NewText(),
	}

	a.WriteString(fmt.Sprintf("<a name=\"%s\"></a>\n", text))

	return a
}
