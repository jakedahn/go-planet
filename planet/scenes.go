package planet

import (
	"net/http"

	"github.com/dghubble/sling"
)

type SceneService struct {
	sling *sling.Sling
}

func NewSceneService(sling *sling.Sling) *SceneService {
	return &SceneService{
		sling: sling.Path("/v0/scenes/ortho"),
	}
}

func (s *SceneService) List(params *FeatureListParams) (*FeatureCollection, *http.Response, error) {
	featureCollection := new(FeatureCollection)
	resp, err := s.sling.New().QueryStruct(params).Receive(featureCollection, nil)

	if err != nil {
		return featureCollection, resp, err
	}

	return featureCollection, resp, nil
}
