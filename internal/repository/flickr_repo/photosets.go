package flickr_repo

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2"
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

func (f *FlickrRepository) GetPhotoSets(userId string, page int) (*photosets.PhotosetsListResponse, error) {
	response, err := photosets.GetList(f.Client, true, userId, page)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (f *FlickrRepository) AddPhotosToPhotoset(photosetId string, photoIds []string) (bool, error) {
	for index, photoId := range photoIds {
		ok, err := addPhotoToPhotoset(f.Client, photosetId, photoId, index)
		if !ok {
			return false, err
		}
	}
	return true, nil
}
func addPhotoToPhotoset(client *flickr.FlickrClient, photosetId string, photoId string, index int) (bool, error) {
	_, err := photosets.AddPhoto(client, photosetId, photoId)
	if err != nil {
		logrus.Errorf("Error while adding a photo with id %s to photoset id %s at index %d: %v", photoId, photosetId, index, err)
		return false, err
	}
	return true, nil
}

func (f *FlickrRepository) RemovePhotosFromPhotoset(photosetId string, photoIds []string) (bool, error) {
	for index, photoId := range photoIds {
		ok, err := removePhotoFromPhotoset(f.Client, photosetId, photoId, index)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

func removePhotoFromPhotoset(client *flickr.FlickrClient, photosetId string, photoId string, index int) (bool, error) {
	_, err := photosets.RemovePhoto(client, photosetId, photoId)
	if err != nil {
		logrus.Errorf("Error while removing a photo with id %s to photoset id %s at index %d: %v", photoId, photosetId, index, err)
		return false, err
	}
	return true, nil
}
