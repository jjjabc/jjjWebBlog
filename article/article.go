package article

import (
	"errors"
	"fmt"
	"jjjBlog/orm"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type JJJarticle struct {
	Title         string    `json:"Title"`
	Id            int       `json:"Id"`
	Text          string    `json:"Text"`
	Imgurl        string    `json:"Imgurl"`
	PublishedTime time.Time `json:"PublishedTime"`
	IsPublished   bool      `json:"IsPublished"`
	Category      string    `json:"Category"`
}

func (this *JJJarticle) AddArticle() error {
	jaId, err := redis.Int(orm.Red.Do("INCR", "art:count"))
	if err != nil {
		return err
	}
	this.Id = jaId
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":title", this.Title)
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":text", this.Text)
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":img", this.Imgurl)
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":cg", this.Category)
	orm.Red.Send("SADD", "art:IdSets", strconv.Itoa(jaId))
	orm.Red.Send("SADD", "Sets:"+this.Category, strconv.Itoa(jaId))
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
func (this *JJJarticle) DelArticle() error {
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":title")
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":text")
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":img")
	orm.Red.Send("DEL", "art:"+strconv.Itoa(this.Id)+":cg")
	orm.Red.Send("SREM", "art:IdSets", strconv.Itoa(this.Id))
	orm.Red.Send("SREM", "Sets:"+this.Category, strconv.Itoa(this.Id))
	if err := orm.Red.Flush(); err != nil {
		return err
	}

	return this.UnPublish()
}
func (this *JJJarticle) UpdataArticle() error {
	if this.Id == 0 {
		return errors.New("Id is zero")
	}
	jaId := this.Id
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":title", this.Title)
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":text", this.Text)
	orm.Red.Send("SET", "art:"+strconv.Itoa(jaId)+":img", this.Imgurl)
	return orm.Red.Flush()
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
	ja.Category, _ = redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+":cg"))
	ja.IsPublished, _ = redis.Bool(orm.Red.Do("SISMEMBER", "art:publishedSets", strconv.Itoa(ArticleId)))
	timeString, _ := redis.String(orm.Red.Do("GET", "art:"+strconv.Itoa(ArticleId)+"publishedTime"))
	dbpubtime, _ := time.Parse(timeString, timeString)
	ja.PublishedTime = dbpubtime
	return &ja
}

func GetPublishedArticlesByCategory(pageNum int, number int, category string) ([]JJJarticle, error) {
	return getArticlesByCategory(pageNum, number, true, category)
}

func GetPublishedArticles(pageNum int, number int) ([]JJJarticle, error) {
	return getArticles(pageNum, number, true)
}
func GetAllArticles() ([]JJJarticle, error) {
	//2147483647是32位系统中Int的最大值
	return getArticles(1, 2000000000, false)
}
func getArtsId(isPublished bool) ([]string, error) {
	var Sets string
	if isPublished {
		Sets = "art:publishedSets"
	} else {
		Sets = "art:IdSets"
	}
	return redis.Strings(orm.Red.Do("SMEMBERS", Sets))
}
func getArtsIdByCategory(isPublished bool, category string) ([]string, error) {
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
func getArticlesByCategory(pageNum int, number int, isPublished bool, category string) ([]JJJarticle, error) {
	all, err := getArtsIdByCategory(isPublished, category)
	if len(all) == 0 {
		return make([]JJJarticle, 0), nil
	}
	if err != nil {
		return nil, errors.New("DB error")
	}
	jaSets := make([]JJJarticle, 0)
	start := (pageNum - 1) * number
	last := len(all) - 1

	for i := start; (i < (start + number)) && (i <= last); i++ {
		aId, _ := strconv.Atoi(all[i])
		ja := GetArticle(aId)
		jaSets = append(jaSets, *ja)
	}
	return jaSets, nil
}

func getArticles(pageNum int, number int, isPublished bool) ([]JJJarticle, error) {
	all, err := getArtsId(isPublished)
	if len(all) == 0 {
		return make([]JJJarticle, 0), nil
	}
	if err != nil {
		return nil, errors.New("DB error")
	}
	jaSets := make([]JJJarticle, 0)
	start := (pageNum - 1) * number
	last := len(all) - 1

	for i := start; (i < (start + number)) && (i <= last); i++ {
		aId, _ := strconv.Atoi(all[i])
		ja := GetArticle(aId)
		jaSets = append(jaSets, *ja)
	}
	return jaSets, nil

}

//获取相对于当前文章的上一篇和下一篇文章
//返回：上一篇artId和下一篇artId，如果没有文章了返回0
func (this *JJJarticle) GetRoundId() (int, int) {
	all, err := getArtsId(this.IsPublished)
	preId := 0
	nextId := 0
	if this.Id == 0 {
		return 0, 0
	}
	thisIdstr := strconv.Itoa(this.Id)
	//没有文章 或 只有一篇文章
	if (len(all) == 0) || (len(all) == 1) {
		return 0, 0
	}

	//第一篇文章
	if thisIdstr == all[0] {
		if nextId, err = strconv.Atoi(all[1]); err != nil {
			return 0, 0
		}
		return 0, nextId
	}

	//最后一篇文章
	if thisIdstr == all[len(all)-1] {
		if preId, err = strconv.Atoi(all[len(all)-2]); err != nil {
			return 0, 0
		}
		return preId, 0
	}

	//其余文章(第2篇——倒数第二篇)
	for i := 1; i < (len(all) - 1); i++ {
		if thisIdstr == all[i] {
			preId, err = strconv.Atoi(all[i-1])
			nextId, err = strconv.Atoi(all[i+1])
			if err != nil {
				return 0, 0
			}
			break
		}
	}
	return preId, nextId
}
