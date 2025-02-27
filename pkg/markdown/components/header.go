package components

type Header struct {
	*Text
}

func NewHeader(level int, str string) *Header {
	h := &Header{
		Text: NewText(),
	}

	h.WriteString("#")
	for i := 0; i < level; i++ {
		h.WriteString("#")
	}

	h.WriteString(" " + str + "\n")

	return h
}
