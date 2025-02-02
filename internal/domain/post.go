package domain

type Post struct {
	ID        uint
	Title     string
	Content   string
	FileID    uint
	SiteID    uint
	CreatedAt string
	UpdatedAt string
}

type CreatePost struct {
	Title   string
	Content string
	FileID  uint
	SiteID  uint
}

type UpdatePost struct {
	Title   string
	Content string
	FileID  uint
}

type PostCommand interface {
	CreatePost(post CreatePost) Post
	UpdatePost(id uint, post UpdatePost) Post
	DeletePost(id uint) Post
}
