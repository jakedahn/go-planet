package planet

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("NewClient gives a pointer to a sling client", t, func() {
		client := NewClient("testkey")
		So(client.sling, ShouldNotBeNil)
		So(client.apiKey, ShouldEqual, "testkey")
	})
}
