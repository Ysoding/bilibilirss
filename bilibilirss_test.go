package bilibilirss_test

import (
	"testing"

	"github.com/Ysoding/bilibilirss"
)

func TestBasic(t *testing.T) {
	c := bilibilirss.New("", "")
	rss, _ := c.GetUpLikeVideo("208259")
	t.Log(rss)
}
