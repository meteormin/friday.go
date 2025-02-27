package components

const CodeBlock = "```"

type QuotingCode struct {
	*Text
	lang string
}

func (qc *QuotingCode) Lang(lang string) {
	qc.lang = lang
}

func NewQuotingCode(str string) *QuotingCode {
	qc := &QuotingCode{
		Text: NewText(),
		lang: "",
	}

	qc.WriteString(str)

	return qc
}

func (qc *QuotingCode) String() string {
	qc.InsertString(CodeBlock + " " + qc.lang + "\n")
	qc.WriteString(CodeBlock)

	return qc.Text.String()
}

type QuotingText struct {
	*Text
}

func NewQuotingText(str string) *QuotingText {
	qc := &QuotingText{
		Text: NewText(),
	}

	qc.WriteString("> ")
	qc.WriteString(str)
	return qc
}
