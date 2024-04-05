package container

import (
	"container/list"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	s := &stack[string]{
		l: list.New(),
	}

	Convey("Stack operations", t, func() {
		Convey("Pop a nil stack", func() {
			Convey("value should be nil", func() {
				So(s.Pop(), ShouldNotBeNil)
			})
		})
		Convey("Pop HelloWorld", func() {
			s.Push("HelloWorld")
			So(s.Pop(), ShouldEqual, "HelloWorld")
		})
		Convey("Length", func() {
			So(s.Len(), ShouldEqual, 0)
		})
	})
}
