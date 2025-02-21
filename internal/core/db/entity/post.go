package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	Path    string
	FileID  uint
	File    *File
	SiteID  uint
	Site    *Site
	Tags    []Tag
}

type Tag struct {
	gorm.Model
	PostID uint
	Post   *Post
	Tag    string
}
