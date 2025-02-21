package domain

import "time"

type Site struct {
	ID        uint
	Host      string
	Name      string
	User      User
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (site *Site) Update(updateSite Site) {
	site.Name = updateSite.Name
	site.UpdatedAt = time.Now()
}
