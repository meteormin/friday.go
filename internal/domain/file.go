package domain

type File struct {
	ID         uint
	OriginName string
	Extension  string
	Size       uint
	FilePath   string
}
