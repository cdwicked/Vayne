package Vayne

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewVayne(t *testing.T) {
	Convey("NewVayne", t, func() {
		Convey("success", func() {
			nv := NewVayne("./mock/file.yaml", "./mock/file.yaml")
			So(nv, ShouldNotBeNil)
		})
	})
}
