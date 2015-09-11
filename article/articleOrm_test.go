package article

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jjjabc/jjjWebBlog/orm"
)
func init(){
	err:=orm.Init("127.0.0.1","6379","")
	if err!=nil{
		panic("redis init error")
	}
}
func TestAddFunc(t *testing.T){
	Convey("测试AddArticleToRedis",t, func() {
		ja1:=JJJarticle{Title:"TestAddArticleToRedis-Title-1",Text:"Text Content 1"}
		err1:=AddArticleToRedis(ja1)
		So(err1,ShouldBeNil)
		So(ja1.Id,ShouldNotEqual,0)
		ja2:=JJJarticle{}
		err2:=AddArticleToRedis(ja2)
		So(err2,ShouldBeNil)
	})
}
func TestGetFunc(t *testing.T)  {
	Convey("测试GetArticle",t, func() {
		ja1:=JJJarticle{Title:"TestGetArticleToRedis-Title-1",Text:"Text Content 1"}
		if err:=AddArticleToRedis(ja1);err!=nil{
			panic("TestGetFunc add article error")
		}
		result:=GetArticle(ja1.Id)
		So(result,ShouldEqual,ja1)
	})
}