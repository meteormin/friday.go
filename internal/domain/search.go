package domain

func FindPostIdByTag(tags []string) map[string][]uint {
	m := make(map[string][]uint)
	for _, t := range tags {
		m[t] = GetCache().Get(t)
	}
	return m
}

func SavePostIdByTag(postId uint, tags []string) {
	for _, t := range tags {
		GetCache().Add(t, postId)
	}
}
