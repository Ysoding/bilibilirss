package bilibilirss_test

import (
	"testing"

	"github.com/Ysoding/bilibilirss"
)

func TestBasic(t *testing.T) {
	c := bilibilirss.New("", "")
	rss, _ := c.GetLeaderboard(0, 3, 1)
	t.Log(rss)
}
