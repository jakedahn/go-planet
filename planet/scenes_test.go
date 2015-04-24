package planet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlanetScenes(t *testing.T) {

	Convey("Test Scenes Client", t, func() {
		mux, testServer := testServer()
		defer testServer.Close()

		Convey("Test Scene List /v0/scenes/ortho", func() {
			file, err := ioutil.ReadFile("../test/scenes_index.json")

			stuff := new(FeatureCollection)
			err = json.Unmarshal(file, stuff)
			So(err, ShouldBeNil)

			req := &http.Request{}
			mux.HandleFunc("/v0/scenes/ortho", func(w http.ResponseWriter, r *http.Request) {
				req = r
				fmt.Print(file)
			})

			fmt.Println("###")
			fmt.Println(req)
			fmt.Println("###")

			client := NewClient("testkey", testServer.URL)
			params := &FeatureListParams{SatId: "1234"}
			featureCollection, rsp, err := client.Scenes.List(params)

			So(err, ShouldBeNil)

			So(rsp, ShouldNotBeNil)
			So(featureCollection, ShouldNotBeNil)
		})

	})
}
