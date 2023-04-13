package flickr_repo

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2"
	"os"
)

type FlickrRepository struct {
	Client *flickr.FlickrClient
}

func (f *FlickrRepository) Authen() (*flickr.FlickrClient, string) {
	apiKey := os.Getenv("FLICKR_API_KEY")
	apiKeySecret := os.Getenv("FLICKR_API_SECRET")
	client := flickr.NewFlickrClient(apiKey, apiKeySecret)

	// first, get a request token
	requestTok, _ := flickr.GetRequestToken(client)

	// build the authorizatin URL
	url, _ := flickr.GetAuthorizeUrl(client, requestTok)

	// get oauth_verifier param in callback url
	accessTok, err := flickr.GetAccessToken(client, requestTok, "<oauth_confirmation_code>")
	if err != nil {
		logrus.Errorf("Error while getting access token: %v", err)
	}
	client.OAuthToken = accessTok.OAuthToken
	client.OAuthTokenSecret = accessTok.OAuthTokenSecret
	return client, url
}
