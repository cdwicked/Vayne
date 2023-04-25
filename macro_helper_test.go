package Vayne

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetMap(t *testing.T) {
	Convey("GetMap", t, func() {
		Convey("nil", func() {
			a := GetMap()
			So(a, ShouldBeNil)
		})
		Convey("not nil", func() {
			e := readConfig("./mock/file.yaml")
			So(e, ShouldBeNil)
			a := GetMap()
			So(a, ShouldNotBeNil)
		})
	})
}

func Test_readConfig(t *testing.T) {
	Convey("readConfig", t, func() {
		Convey("success", func() {
			a := readConfig("./mock/file.yaml")
			So(a, ShouldBeNil)
		})
	})
}
