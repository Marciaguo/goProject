package bdd

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName(t *testing.T) {
	convey.Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4
		convey.Convey("when add the two numbers", func() {
			c := a + b
			convey.Convey("when the result is still even", func() {
				convey.So(c%2, convey.ShouldEqual, 0)
			})
		})
	})
}
