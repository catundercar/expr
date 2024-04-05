package container

import (
	"container/list"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	s := &queue[string]{
		l: list.New(),
	}

	Convey("Stack operations", t, func() {
		Convey("Pop a nil stack", func() {
			Convey("value should be nil", func() {
				So(s.Pop(), ShouldEqual, "")
			})
		})
		Convey("Pop HelloWorld", func() {
			s.Push("HelloWorld1")
			s.Push("HelloWorld2")
			So(s.Pop(), ShouldEqual, "HelloWorld1")
			So(s.Pop(), ShouldEqual, "HelloWorld2")
		})
		Convey("Length", func() {
			So(s.Len(), ShouldEqual, 0)
		})
	})
}
