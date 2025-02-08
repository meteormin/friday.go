package domain

type Post struct {
	ID        uint
	Title     string
	Content   string
	FileID    uint
	File      *File
	SiteID    uint
	Site      *Site
	Tag       []Tag
	CreatedAt string
	UpdatedAt string
}

type Site struct {
	ID    uint
	Host  string
	Alias string
}

type Tag struct {
	ID  uint
	Tag string
}
