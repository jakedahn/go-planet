package planet

import (
	"net/http"

	sling "github.com/dghubble/sling"
)

type Client struct {
	sling  *sling.Sling
	Scenes *SceneService

	apiKey     string
	BaseUrl    string
	httpClient *http.Client
}

func NewClient(apiKey string, baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = "https://api.planet.com/v0/"
	}

	httpClient := &http.Client{}
	base := sling.New().Client(httpClient).Base(baseUrl)

	return &Client{
		sling:      base,
		apiKey:     apiKey,
		httpClient: httpClient,

		Scenes: NewSceneService(base.New()),
	}
}
