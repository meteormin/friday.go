package component

type Comment struct {
	*Text
}

func NewComment(str string) *Comment {
	c := &Comment{
		Text: NewText(),
	}

	c.WriteString("<!-- ")
	c.WriteString(str)
	c.WriteString(" -->")

	return c
}
