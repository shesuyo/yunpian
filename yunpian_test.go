package yunpian

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var yourapikey = ""

func TestRegisterKey(t *testing.T) {
	Convey("TestRegisterKey", t, func() {
		RegisterKey(yourapikey)
		So(APIKEY, ShouldEqual, yourapikey)
	})

}

func TestUserGet(t *testing.T) {
	Convey("UserGet BeNil", t, func() {
		RegisterKey(yourapikey)
		_, err := UserGet()
		So(err, ShouldBeNil)
	})
	Convey("UserGet NotNil", t, func() {
		RegisterKey("")
		_, err := UserGet()
		So(err, ShouldNotBeNil)
	})
}

func TestTplGet(t *testing.T) {
	Convey("TestTplGet BeNil", t, func() {
		RegisterKey(yourapikey)
		_, err := Tpl_Get(1)
		So(err, ShouldBeNil)
	})
}
