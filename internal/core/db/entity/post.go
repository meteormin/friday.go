package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	FileID  uint
	File    *File
	SiteID  uint
	Site    *Site
	Tag     []Tag
}

type Tag struct {
	gorm.Model
	PostID uint
	Post   *Post
	Tag    string
}
