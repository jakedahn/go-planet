package planet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlanetScenes(t *testing.T) {
	Convey("Given a valid Scenes Client", t, func() {
		// Test setup
		jsonFixture, err := ioutil.ReadFile("../test/scenes_index.json")
		if err != nil {
			fmt.Errorf("Error: %s", err)
		}
		mux, testServer := testServer()
		defer testServer.Close()

		// Setting up the test http server endpoint
		testServerRequest := &http.Request{}
		mux.HandleFunc("/v0/scenes/ortho", func(w http.ResponseWriter, r *http.Request) {
			testServerRequest = r
			w.Write(jsonFixture)
		})

		// instantiating new client
		client := NewClient("testkey", testServer.URL)
		So(client.apiKey, ShouldEqual, "testkey")
		So(client.sling.RawUrl, ShouldEqual, testServer.URL)

		Convey("the List method should", func() {
			params := &FeatureListParams{SatId: "0815", FileSizeGreaterThan: "10"}
			featureCollection, rsp, err := client.Scenes.List(params)
			So(err, ShouldBeNil)

			Convey("parse and send the given FeatureListParams", func() {
				expectedParams := "file_size.gt=10&sat.id=0815"
				So(testServerRequest.URL.RawQuery, ShouldEqual, expectedParams)
				So(rsp.StatusCode, ShouldEqual, 200)
			})

			Convey("initiate a call to the correct URL", func() {
				So(testServerRequest.URL.Path, ShouldEqual, "/v0/scenes/ortho")
				So(rsp.StatusCode, ShouldEqual, 200)
			})

			Convey("call the API and parse the resulting JSON into a FeatureCollection", func() {
				So(len(featureCollection.Features), ShouldEqual, 3)
				So(featureCollection.Features[0].Id, ShouldEqual, "test_id_1")
			})

		})

	})
}
