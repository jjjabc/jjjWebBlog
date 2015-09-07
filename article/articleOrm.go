package orm

import (
	"github.com/jjjabc/jjjWebBlog/article"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
	"errors"
)
func GetArticle(ArticleId int) *article.JJJarticle {
	ja := article.JJJarticle{
		Id: ArticleId,
	}
	var err error
	ja.Title, err = redis.String(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":title"))
	if err != nil {
		return nil
	}
	ja.Text, _ = redis.String(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":text"))
	ja.Imgurl, _ = redis.String(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":img"))
	ja.Category, _ = redis.String(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":cg"))
	ja.IsPublished, _ = redis.Bool(Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(ArticleId)))
	timeString, _ := redis.String(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + "publishedTime"))
	dbpubtime, _ := time.Parse(timeString, timeString)
	ja.PublishedTime = dbpubtime
	ja.Priority, _ = redis.Int(Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":pri"))
	return &ja
}
func GetArtsId(isPublished bool) ([]string, error) {
	var Sets string
	if isPublished {
		Sets = "art:publishedSets"
	} else {
		Sets = "art:IdSets"
	}
	return redis.Strings(Red.Do("SMEMBERS", Sets))
}
func GetArtsIdByCategory(isPublished bool, category string) ([]string, error) {
	var Sets string
	if isPublished {
		Sets = "art:publishedSets"
	} else {
		Sets = "art:IdSets"
	}
	cgsets := "Sets:" + category
	fmt.Println(cgsets)
	return redis.Strings(Red.Do("SINTER", Sets, cgsets))
}
func AddArticleToRedis(ja article.JJJarticle) error{
	jaId, err := redis.Int(Red.Do("INCR", "art:count"))
	if err != nil {
		return err
	}
	ja.Id = jaId
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":title", ja.Title)
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":text", ja.Text)
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":img", ja.Imgurl)
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":cg", ja.Category)
	Red.Send("SADD", "art:IdSets", strconv.Itoa(jaId))
	Red.Send("SADD", "Sets:" + ja.Category, strconv.Itoa(jaId))
	return Red.Flush()
}
func DelArticleFromRedis(ja article.JJJarticle)error{
	Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":title")
	Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":text")
	Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":img")
	Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":cg")
	Red.Send("SREM", "art:IdSets", strconv.Itoa(ja.Id))
	Red.Send("SREM", "Sets:" + ja.Category, strconv.Itoa(ja.Id))
	if err := Red.Flush(); err != nil {
		return err
	}
	return ja.UnPublish()
}

//true:已发布 false:未发布或id不存在
func PublishStatus(id int)bool{
	if exist, _ := redis.Int(Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(id))); exist == 1 {
		return true
	}
	return false
}

//将id号添加到已发布文章集合中
func AddIdToPublishedSet(id int)error{
	_,err:=Red.Do("SADD", "art:publishedSets", strconv.Itoa(id))
	return err
}
func DelIdFromPublishedSet(id int)error{
	_, err := Red.Do("SREM", "art:publishedSets", strconv.Itoa(id))
	return err

}

func SetPublishTime(ja article.JJJarticle)error{
	_,err:=Red.Do("SET", "art:" + strconv.Itoa(ja.Id) + "publishedTime", ja.PublishedTime.String())
	return err
}
func UpdateArticle(ja article.JJJarticle)error{
	if ja.Id == 0 {
		return errors.New("Id is zero")
	}
	jaId := ja.Id
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":title", ja.Title)
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":text", ja.Text)
	Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":img", ja.Imgurl)
	return Red.Flush()
}