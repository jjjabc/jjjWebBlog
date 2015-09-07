package article

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
	"errors"
	"github.com/jjjabc/jjjWebBlog/orm"
)
func GetArticle(ArticleId int) *JJJarticle {
	ja := JJJarticle{
		Id: ArticleId,
	}
	var err error
	ja.Title, err = redis.String(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":title"))
	if err != nil {
		return nil
	}
	ja.Text, _ = redis.String(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":text"))
	ja.Imgurl, _ = redis.String(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":img"))
	ja.Category, _ = redis.String(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":cg"))
	ja.IsPublished, _ = redis.Bool(orm.Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(ArticleId)))
	timeString, _ := redis.String(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + "publishedTime"))
	dbpubtime, _ := time.Parse(timeString, timeString)
	ja.PublishedTime = dbpubtime
	ja.Priority, _ = redis.Int(orm.Red.Do("GET", "art:" + strconv.Itoa(ArticleId) + ":pri"))
	return &ja
}
func GetArtsId(isPublished bool) ([]string, error) {
	var Sets string
	if isPublished {
		Sets = "art:publishedSets"
	} else {
		Sets = "art:IdSets"
	}
	return redis.Strings(orm.Red.Do("SMEMBERS", Sets))
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
	return redis.Strings(orm.Red.Do("SINTER", Sets, cgsets))
}
func AddArticleToRedis(ja JJJarticle) error{
	jaId, err := redis.Int(orm.Red.Do("INCR", "art:count"))
	if err != nil {
		return err
	}
	ja.Id = jaId
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":title", ja.Title)
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":text", ja.Text)
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":img", ja.Imgurl)
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":cg", ja.Category)
	orm.Red.Send("SADD", "art:IdSets", strconv.Itoa(jaId))
	orm.Red.Send("SADD", "Sets:" + ja.Category, strconv.Itoa(jaId))
	return orm.Red.Flush()
}
func DelArticleFromRedis(ja JJJarticle)error{
	orm.Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":title")
	orm.Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":text")
	orm.Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":img")
	orm.Red.Send("DEL", "art:" + strconv.Itoa(ja.Id) + ":cg")
	orm.Red.Send("SREM", "art:IdSets", strconv.Itoa(ja.Id))
	orm.Red.Send("SREM", "Sets:" + ja.Category, strconv.Itoa(ja.Id))
	if err := orm.Red.Flush(); err != nil {
		return err
	}
	return ja.UnPublish()
}

//true:已发布 false:未发布或id不存在
func PublishStatus(id int)bool{
	if exist, _ := redis.Int(orm.Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(id))); exist == 1 {
		return true
	}
	return false
}

//将id号添加到已发布文章集合中
func AddIdToPublishedSet(id int)error{
	_,err:=orm.Red.Do("SADD", "art:publishedSets", strconv.Itoa(id))
	return err
}
func DelIdFromPublishedSet(id int)error{
	_, err := orm.Red.Do("SREM", "art:publishedSets", strconv.Itoa(id))
	return err

}

func SetPublishTime(ja JJJarticle)error{
	_,err:=orm.Red.Do("SET", "art:" + strconv.Itoa(ja.Id) + "publishedTime", ja.PublishedTime.String())
	return err
}
func UpdateArticle(ja JJJarticle)error{
	if ja.Id == 0 {
		return errors.New("Id is zero")
	}
	jaId := ja.Id
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":title", ja.Title)
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":text", ja.Text)
	orm.Red.Send("SET", "art:" + strconv.Itoa(jaId) + ":img", ja.Imgurl)
	return orm.Red.Flush()
}