package http

import "time"

type ContentResource[T interface{}] struct {
	Content []T `json:"content"`
	Length  int `json:"length"`
}

func NewContentResource[T interface{}](content []T) ContentResource[T] {
	return ContentResource[T]{
		Content: content,
		Length:  len(content),
	}
}

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format("2006-01-02 15:04:05") + `"`), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", string(data))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func NewDateTime(t time.Time) *DateTime {
	if t.IsZero() {
		return nil
	}

	return &DateTime{
		Time: t,
	}
}
