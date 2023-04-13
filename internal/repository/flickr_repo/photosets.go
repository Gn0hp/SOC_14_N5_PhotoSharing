package flickr_repo

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2/photosets"
)

func (f *FlickrRepository) CreatePhotoset(title, description, primaryPhotoID string) (photosets.PhotosetResponse, error) {
	response, err := photosets.Create(f.Client, title, description, primaryPhotoID)
	if err != nil {
		logrus.Errorf("Error while creating photoset: %v", err)
		return photosets.PhotosetResponse{}, err
	}
	return *response, nil
}

func (f *FlickrRepository) GetPhotoSets() {

}
