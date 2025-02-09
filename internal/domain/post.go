package domain

import "time"

type Post struct {
	ID        uint
	Title     string
	Content   string
	Path      string
	FileID    uint
	File      *File
	SiteID    uint
	Site      *Site
	Tags      []Tag
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Site struct {
	ID    uint
	Host  string
	Name  string
	Posts []Post
}

type Tag struct {
	ID  uint
	Tag string
}
