package domain

import (
	"github.com/meteormin/friday.go/pkg/cache"
	"time"
)

func FindPostIdByTag(tags []string) map[string][]uint {
	m := make(map[string][]uint)
	for _, t := range tags {
		var postId uint
		cache.Get(t, &postId)
		m[t] = append(m[t], postId)
	}
	return m
}

func SavePostIdByTag(postId uint, exp time.Duration, tags []string) error {
	for _, t := range tags {
		err := cache.Set(t, exp, postId)
		if err != nil {
			return err
		}
	}
	return nil
}
