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
		sling: sling.Path("/v0/scenes/ortho/"),
	}
}

func (s *SceneService) List(params *FeatureListParams) (*FeatureCollection, *http.Response, error) {
	featureCollection := &FeatureCollection{}
	resp, err := s.sling.New().Get("").Receive(featureCollection)
	if err != nil {
		return featureCollection, resp, err
	}

	return featureCollection, resp, nil
}
