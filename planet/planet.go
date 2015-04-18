package planet

import (
	"net/http"

	sling "github.com/dghubble/sling"
)

const apiBaseUrl = "https://api.planet.com/v0/"

type Client struct {
	sling  *sling.Sling
	apiKey string
}

func NewClient(apiKey string) *Client {
	httpClient := &http.Client{}
	base := sling.New().Client(httpClient).Base(apiBaseUrl)

	return &Client{
		sling:  base,
		apiKey: apiKey,
	}
}
