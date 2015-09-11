package article

import (
	"errors"
	"strconv"
	"time"
	"sort"
)

type JJJarticle struct {
	Title         string    `json:"Title"`
	Id            int       `json:"Id"`
	Text          string    `json:"Text"`
	Imgurl        string    `json:"Imgurl"`        //标题图片
	PublishedTime time.Time `json:"PublishedTime"` //发布时间
	IsPublished   bool      `json:"IsPublished"`   //是否发布，ture：已发布，false：未发布
	Category      string    `json:"Category"`      //所属栏目
	Priority      int       `json:"Priority"`     //优先级，优先级越高文章显示越前面。Priority小于0：置顶，优先级相同已时间排序为准。默认优先级为0
}

func (this *JJJarticle) AddArticle() error {
	return AddArticleToRedis(*this)
}

//发布文章，redis中使用Sets来保存已发布的文章ID
func (this *JJJarticle) Publish() error {
	AlreadyPubErr := errors.New("Article is published!!")
	if this.IsPublished == true {
		return AlreadyPubErr
	}
	if PublishStatus(this.Id) {
		this.IsPublished=true
		return AlreadyPubErr
	}

	if err := AddIdToPublishedSet(this.Id); err != nil {
		return err
	}

	this.PublishedTime = time.Now()
	if err := SetPublishTime(*this); err != nil {
		return err
	}
	this.IsPublished = true
	return nil
}
func (this *JJJarticle) UnPublish() error {
	this.IsPublished = false
	return DelIdFromPublishedSet(this.Id)
}
func (this *JJJarticle) DelArticle() error {
	return DelArticleFromRedis(*this)
}
func (this *JJJarticle) UpdataArticle() error {
	if this.Id == 0 {
		return errors.New("Id is zero")
	}
	return UpdateArticle(*this)
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

func getArticlesByCategory(pageNum int, number int, isPublished bool, category string) ([]JJJarticle, error) {
	all, err := GetArtsIdByCategory(isPublished, category)
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
	sort.Reverse(ArticleSortSlice(jaSets))
	return jaSets, nil
}

func getArticles(pageNum int, number int, isPublished bool) ([]JJJarticle, error) {
	all, err := GetArtsId(isPublished)
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
	sort.Reverse(ArticleSortSlice(jaSets))
	return jaSets, nil
}

//获取相对于当前文章的上一篇和下一篇文章
//返回：上一篇artId和下一篇artId，如果没有文章了返回0
func (this *JJJarticle) GetRoundId() (int, int) {
	all, err := GetArtsId(this.IsPublished)
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
	if thisIdstr == all[len(all) - 1] {
		if preId, err = strconv.Atoi(all[len(all) - 2]); err != nil {
			return 0, 0
		}
		return preId, 0
	}

	//其余文章(第2篇——倒数第二篇)
	for i := 1; i < (len(all) - 1); i++ {
		if thisIdstr == all[i] {
			preId, err = strconv.Atoi(all[i - 1])
			nextId, err = strconv.Atoi(all[i + 1])
			if err != nil {
				return 0, 0
			}
			break
		}
	}
	return preId, nextId
}
