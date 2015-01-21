package article

import (
	"errors"
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

func (this *JJJarticle) AddArticle() error {
	orm.Red.Send("SET", "art:"+strconv.Itoa(this.Id)+":title", this.Title)
	orm.Red.Send("SET", "art:"+strconv.Itoa(this.Id)+":text", this.Text)
	orm.Red.Send("SET", "art:"+strconv.Itoa(this.Id)+":img", this.Imgurl)
	return orm.Red.Flush()
}

//发布文章，redis中使用Sets来保存已发布的文章ID
func (this *JJJarticle) Publish() error {
	AlreadyPubErr := errors.New("Article is published!!")
	if this.IsPublished == true {
		return AlreadyPubErr
	}
	if exist, _ := redis.Int(orm.Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(this.Id))); exist == 1 {
		return AlreadyPubErr
	}

	if _, err := orm.Red.Do("SADD", "art:publishedSets", strconv.Itoa(this.Id)); err != nil {
		return err
	}

	this.PublishedTime = time.Now()
	if _, err := orm.Red.Do("SET", "art:"+strconv.Itoa(this.Id)+"publishedTime", this.PublishedTime.String()); err != nil {
		return err
	}
	this.IsPublished = true
	return nil
}
func (this *JJJarticle) UnPublish() error {
	this.IsPublished = false
	_, err := orm.Red.Do("SREM", "art:publishedSets", strconv.Itoa(this.Id))
	return err
}
func (this *JJJarticle) DelArticle() {
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":title")
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":text")
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":img")
	this.UnPublish()
	return
}
func GetArticle(ArticleId int) *JJJarticle {
	ja := JJJarticle{
		Id: ArticleId,
	}
	var err error
	ja.Title, err = redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+":title"))
	if err != nil {
		return nil
	}
	ja.Text, _ = redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+":text"))
	ja.Imgurl, _ = redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+":img"))
	ja.IsPublished, _ = redis.Bool(orm.Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(ArticleId)))
	timeString, _ := redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+"publishedTime"))
	dbpubtime, _ := time.Parse(timeString, timeString)
	ja.PublishedTime = dbpubtime
	return &ja
}
func GetPublishedArticles()
func GetAllArticles()
