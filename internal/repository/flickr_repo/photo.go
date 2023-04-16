package flickr_repo

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/pkg/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2"
	"gopkg.in/masci/flickr.v2/people"
	"gopkg.in/masci/flickr.v2/photos"
	"io"
)

func (f *FlickrRepository) UploadPhoto(reader io.Reader, name string) (*flickr.UploadResponse, bool) {
	response, err := flickr.UploadReader(f.Client, reader, name, nil)
	if err != nil {
		logrus.Errorf("Error while uploadding photo to Flickr: %v", err)
		return nil, false
	}
	return response, true
}

func (f *FlickrRepository) GetPhotoInfo(id, secret string) (*photos.PhotoInfoResponse, error) {
	response, err := photos.GetInfo(f.Client, id, secret)
	if err != nil {
		logrus.Errorf("Error while getting photo info: %v", err)
		return nil, err
	}
	return response, nil
}

func (f *FlickrRepository) GetPhotos(userID string) (*entities.PhotoListResponse, error) {
	response, err := people.GetPhotos(f.Client, userID, people.GetPhotosOptionalArgs{
		SafeSearch:    0,
		MinUploadDate: "",
		MaxUploadDate: "",
		MinTakenDate:  "",
		MaxTakenDate:  "",
		ContentType:   0,
		PrivacyFilter: 0,
		Extras:        "url_sq, url_t, url_s, url_q, url_m, url_n, url_z, url_c, url_l, url_o",
		PerPage:       0,
		Page:          0,
	})
	if err != nil {
		logrus.Errorf("Error while getting people photo: %v", err)
		return nil, err
	}
	return utils.ParsePhotosXmlToPhotoListResponse(response.Extra), nil
}
