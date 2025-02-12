package component

import (
	"net/url"
)

type Link struct {
	*Text
}

func NewLink(text string, url url.URL) *Link {
	l := &Link{
		Text: NewText(),
	}

	l.WriteString("[" + text + "](" + url.String() + ")")

	return l
}

type ImageLink struct {
	*Link
}

func NewImageLink(alt string, url url.URL) *ImageLink {
	il := &ImageLink{
		Link: NewLink(alt, url),
	}

	il.WriteString("!" + il.Link.String() + " " + alt)

	return il
}
