package article

import (
	"github.com/garyburd/redigo/redis"
	"jjjBlog/orm"
	"strconv"
	"time"
)

type JJJarticle struct {
	Title         string
	Id            int
	Text          string
	Imgurl        string
	PublishedTime time.Time
	IsPublished   bool
}

func AddArticle(ja JJJarticle) error {
	orm.Red.Send("SET", "artDraft:"+strconv.Itoa(ja.Id)+":title", ja.Title)
	orm.Red.Send("SET", "artDraft:"+strconv.Itoa(ja.Id)+":text", ja.Text)
	orm.Red.Send("SET", "artDraft:"+strconv.Itoa(ja.Id)+":img", ja.Imgurl)
}
