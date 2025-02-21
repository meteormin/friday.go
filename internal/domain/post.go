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

func (post *Post) Update(updatePost Post) {
	post.Title = updatePost.Title
	post.Content = updatePost.Content
	post.Path = updatePost.Path
	post.FileID = updatePost.FileID
	post.UpdatedAt = time.Now()
}

type Tag struct {
	ID  uint
	Tag string
}
