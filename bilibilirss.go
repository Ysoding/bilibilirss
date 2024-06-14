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

var RID2ChineseNameMap = map[int]string{
	0:   "全站",
	1:   "动画",
	168: "国创相关",
	3:   "音乐",
	129: "舞蹈",
	4:   "游戏",
	36:  "科技",
	188: "数码",
	160: "生活",
	119: "鬼畜",
	155: "时尚",
	5:   "娱乐",
	181: "影视",
}

func New(cookie string, uid string) *BiliBiliRss {
	return &BiliBiliRss{
		Cookie: cookie,
		UID:    uid,
		client: *collect.NewClient(cookie, uid),
	}
}

// GetLeaderboard 获取排行榜信息
//
// Parameters:
// - rid: 排行榜分区 id ( 全站(0)  动画(1)  国创相关(168)  音乐(3)  舞蹈(129)  游戏(4)  科技(36)  数码(188) 生活(160) 鬼畜(119)  时尚(155)  娱乐(5)  影视(181) )
// - day: 时间跨度, 可为 1 3 7 30
// - arcType: 投稿时间, 可为 0(全部投稿) 1(近期投稿)
func (b *BiliBiliRss) GetLeaderboard(rid, day, arcType int) (string, error) {
	body, err := b.client.GetLeaderboard(rid, day, arcType)
	if err != nil {
		return "", err
	}

	if gjson.GetBytes(body, "code").Int() != 0 {
		return "", fmt.Errorf("fetch error response %s", string(body))
	}

	data := []rss.RssData{}

	gjson.GetBytes(body, "data.list").ForEach(func(key, value gjson.Result) bool {
		d := &payload.UpLikeVideo{
			Title:  value.Get("title").String(),
			Author: value.Get("author").String(),
			Link:   fmt.Sprintf("https://www.bilibili.com/video/%s", value.Get("bvid").String()),
		}

		if value.Get("create").Exists() {
			d.PubDate = time.Unix(value.Get("create").Int(), 0).Format("2006-01-02 15:04:05")
		}

		data = append(data, d)
		return true
	})

	var archTypeName string
	if arcType == 0 {
		archTypeName = "全部投稿"
	} else {
		archTypeName = "近期投稿"
	}

	r := rss.NewRssRender(data)
	r.RssTitle = fmt.Sprintf("bilibili %d日排行榜-%s-%s", day, RID2ChineseNameMap[rid], archTypeName)
	r.RssLink = fmt.Sprintf("https://www.bilibili.com/ranking/all/%d/0/%d", rid, day)
	return r.Render(), nil

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
	r.RssLink = fmt.Sprintf("https://space.bilibili.com/%s", uid)
	return r.Render(), nil
}
