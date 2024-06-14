package rss

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/Ysoding/bilibilirss/internal/temp"
)

type RssData interface {
	GetItem() Item
}

type Item struct {
	Title   string
	Desc    string
	PubDate string
	Link    string
	Author  string
}

type RssRender struct {
	Items    []Item
	RssNow   template.HTML
	RssTitle string
	RssLink  string
	RssDesc  string
}

func getNow() string {
	return time.Now().Format(time.RFC822Z)
}

func NewRssRender(data []RssData) *RssRender {
	r := &RssRender{
		RssNow: template.HTML(getNow()),
		Items:  make([]Item, 0),
	}

	for _, d := range data {
		r.Items = append(r.Items, d.GetItem())
	}

	return r
}

func (r *RssRender) update() {
	r.RssDesc = fmt.Sprintf("%s - Made with love by Ysoding(https://github.com/Ysoding/bilibilirss)", r.RssTitle)
}

func (r *RssRender) Render() string {
	r.update()
	t, err := template.New("atom").Parse(temp.Tmpl)
	if err != nil {
		log.Fatal(err)
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, r)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
