package article

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_ArticleSortSlice_func(t *testing.T) {
	jas := []JJJarticle{
		JJJarticle{Id: 0, Priority: 10, PublishedTime: time.Unix(1000000, 0)},
		JJJarticle{Id: 1, Priority: -1, PublishedTime: time.Unix(1010000, 0)},
		JJJarticle{Id: 2, Priority: 1, PublishedTime: time.Unix(1020000, 0)},
		JJJarticle{Id: 3, Priority: 10, PublishedTime: time.Unix(1030000, 0)},
		JJJarticle{Id: 4, Priority: -1, PublishedTime: time.Unix(1040000, 0)},
	}
	Convey("测试ArticleSortSlice.Len", t, func() {
		So(ArticleSortSlice(jas).Len(), ShouldEqual, 5)
	})
	Convey("测试ArticleSortSlice.Swap", t, func() {
		ArticleSortSlice(jas).Swap(0, 2)
		So(jas[0].Id, ShouldEqual, 2)
		So(jas[2].Id, ShouldEqual, 0)
		ArticleSortSlice(jas).Swap(0, 2)
	})
	Convey("测试ArticleSortSlice.Less", t, func() {
		Convey("测试ja优先级为－1", func() {
			So(ArticleSortSlice(jas).Less(0, 1), ShouldBeTrue)
			So(ArticleSortSlice(jas).Less(1, 0), ShouldBeFalse)
			So(ArticleSortSlice(jas).Less(1, 4), ShouldBeTrue)
			So(ArticleSortSlice(jas).Less(4, 1), ShouldBeFalse)
		})
		Convey("测试ja优先级>0且优先级不同的情况", func() {
			So(ArticleSortSlice(jas).Less(0, 2), ShouldBeFalse)
			So(ArticleSortSlice(jas).Less(2, 0), ShouldBeTrue)
		})
		Convey("测试ja优先级>0且优先级相同的情况", func() {
			So(ArticleSortSlice(jas).Less(0, 3), ShouldBeTrue)
			So(ArticleSortSlice(jas).Less(3, 0), ShouldBeFalse)
		})
	})
}