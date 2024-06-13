package collect

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	cookie string
	uid    string
}

func NewClient(cookie, uid string) *Client {
	return &Client{
		cookie: cookie,
		uid:    uid,
	}
}

func (c *Client) GetUpLikeVideo(uid string) ([]byte, error) {
	return c.get(fmt.Sprintf("https://api.bilibili.com/x/space/like/video?vmid=%s", uid))
}

func (c *Client) get(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get url failed:%w", err)
	}

	if c.cookie != "" {
		req.Header.Set("Cookie", c.cookie)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)

}
