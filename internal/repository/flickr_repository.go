package repository

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2"
	"gopkg.in/masci/flickr.v2/photosets"
	"io"
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

func (f *FlickrRepository) CreatePhotoset(title, description, primaryPhotoID string) (photosets.PhotosetResponse, error) {
	response, err := photosets.Create(f.Client, title, description, primaryPhotoID)
	if err != nil {
		logrus.Errorf("Error while creating photoset: %v", err)
		return photosets.PhotosetResponse{}, err
	}
	return *response, nil
}

func (f *FlickrRepository) UploadPhoto(reader io.Reader, name string) (*flickr.UploadResponse, bool) {
	response, err := flickr.UploadReader(f.Client, reader, name, nil)
	if err != nil {
		logrus.Errorf("Error while uploadding photo to Flickr: %v", err)
		return nil, false
	}
	return response, true
}
