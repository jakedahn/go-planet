package planet

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testServer() (*http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	return mux, server
}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func TestGet(t *testing.T) {
	Convey("NewClient gives a pointer to a sling client", t, func() {
		client := NewClient("testkey", "")
		So(client.sling, ShouldNotBeNil)
		So(client.apiKey, ShouldEqual, "testkey")
	})
}
