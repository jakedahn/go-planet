package planet

import (
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlanetScenes(t *testing.T) {

	Convey("Test Scenes Client", t, func() {
		jsonFixture, err := ioutil.ReadFile("../test/scenes_index.json")
		So(err, ShouldBeNil)
		mux, testServer := testServer()
		defer testServer.Close()

		Convey("Test Scene List /v0/scenes/ortho", func() {
			req := &http.Request{}
			mux.HandleFunc("/v0/scenes/ortho", func(w http.ResponseWriter, r *http.Request) {
				req = r
				w.Write(jsonFixture)
			})

			client := NewClient("testkey", testServer.URL)
			params := &FeatureListParams{SatId: "0815"}
			featureCollection, rsp, err := client.Scenes.List(params)
			So(err, ShouldBeNil)

			Convey("Test that the server can process all parameters", func() {
				So(req.URL.Path, ShouldEqual, "/v0/scenes/ortho")
			})

			Convey("Test that the client response contains a FeatureCollection", func() {
				So(featureCollection, ShouldNotBeNil)
			})

			So(rsp.StatusCode, ShouldEqual, 200)
		})

	})
}
