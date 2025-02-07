package domain

type File struct {
	ID         uint
	OriginName string
	ConvName   string
	MimeType   string
	Size       uint64
	FilePath   string
}
