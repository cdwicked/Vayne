package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestGetMap(t *testing.T) {
	Convey("GetMap", t, func() {
		Convey("nil", func() {
			a := GetMap("Values")
			So(len(a), ShouldEqual, 0)
		})
		Convey("not nil", func() {
			e := ReadMacroConfig("../mock/values.yaml")
			So(e, ShouldBeNil)
			a := GetMap("Values")
			So(a, ShouldNotBeNil)
			// ${.Values.baseurl}

		})
	})
}

func Test_readConfig(t *testing.T) {
	Convey("readConfig", t, func() {
		Convey("success", func() {
			a := ReadMacroConfig("../mock/values.yaml")
			So(a, ShouldBeNil)
		})
	})
}

func TestReadValue(t *testing.T) {
	Convey("TestReadValue", t, func() {
		Convey("success read multi object", func() {
			err := ReadMacroConfig("../mock/values.yaml")
			So(err, ShouldBeNil)
			str, err := ReadValue(".Values.ccc.bbb")
			log.Printf("result is %v", str)
			So(str, ShouldEqual, "aaa")
		})
		Convey("success read single object", func() {
			err := ReadMacroConfig("../mock/values.yaml")
			So(err, ShouldBeNil)
			str, err := ReadValue(".Values.baseurl")
			log.Printf("result is %v", str)
			So(str, ShouldEqual, "http://1.1.1.1")
		})
		Convey("success read array", func() {
			err := ReadMacroConfig("../mock/values.yaml")
			So(err, ShouldBeNil)
			str, err := ReadValue(".Values.ccc")
			log.Printf("result is %v", str)
			a := make(map[string]interface{}, 0)
			So(str, ShouldHaveSameTypeAs, a)
		})
	})
}
