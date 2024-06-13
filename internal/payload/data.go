package payload

import "github.com/Ysoding/bilibilirss/internal/rss"

type UpLikeVideo struct {
	Title   string
	Desc    string
	PubDate string
	Link    string
	Author  string
}

func (u *UpLikeVideo) GetItem() rss.Item {
	return rss.Item{
		Title:   u.Title,
		Desc:    u.Desc,
		PubDate: u.PubDate,
		Link:    u.Link,
		Author:  u.Link,
	}
}
