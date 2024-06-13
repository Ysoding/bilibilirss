package bilibilirss

import (
	"fmt"
	"time"

	"github.com/Ysoding/bilibilirss/internal/collect"
	"github.com/Ysoding/bilibilirss/internal/payload"
	"github.com/Ysoding/bilibilirss/internal/rss"
	"github.com/tidwall/gjson"
)

type BiliBiliRss struct {
	Cookie string
	UID    string
	client collect.Client
}

func New(cookie string, uid string) *BiliBiliRss {
	return &BiliBiliRss{
		Cookie: cookie,
		UID:    uid,
		client: *collect.NewClient(cookie, uid),
	}
}

// GetUpLike 获取up主点赞的视频
func (b *BiliBiliRss) GetUpLikeVideo(uid string) (string, error) {
	body, err := b.client.GetUpLikeVideo(uid)
	if err != nil {
		return "", err
	}

	if gjson.GetBytes(body, "code").Int() != 0 {
		return "", fmt.Errorf("fetch error response %s", string(body))
	}

	data := []rss.RssData{}
	gjson.GetBytes(body, "data.list").ForEach(func(key, value gjson.Result) bool {
		pubdate := value.Get("pubdate").Int()
		d := &payload.UpLikeVideo{
			Title:   value.Get("title").String(),
			Desc:    value.Get("desc").String(),
			PubDate: time.Unix(pubdate, 0).Format("2006-01-02 15:04:05"),
			Author:  value.Get("owner.name").String(),
			Link:    fmt.Sprintf("https://www.bilibili.com/video/%s", value.Get("bvid").String()),
		}
		data = append(data, d)
		return true
	})

	r := rss.NewRssRender(data)
	r.RssTitle = fmt.Sprintf("%s 的 bilibili 点赞视频", uid)
	r.RssDesc = fmt.Sprintf("%s 的 bilibili 点赞视频 - Made with love by Ysoding(https://github.com/Ysoding/bilibilirss)", uid)
	r.RssLink = fmt.Sprintf("https://space.bilibili.com/%s", uid)
	return r.Render(), nil
}
