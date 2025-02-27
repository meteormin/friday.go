package components

type AlertType string

const (
	NOTE      AlertType = "[!NOTE]]"
	TIP       AlertType = "[!TIP]]"
	WARNING   AlertType = "[!WARNING]]"
	IMPORTANT AlertType = "[!IMPORTANT]]"
)

type Alert struct {
	*Text
	Type AlertType
}

func NewAlert(alertType AlertType, str string) *Alert {
	a := &Alert{
		Type: alertType,
		Text: NewText(),
	}

	qt := NewQuotingText(string(a.Type))
	a.WriteString(qt.String())

	qt = NewQuotingText(str)
	a.WriteString(qt.String())

	return a
}
