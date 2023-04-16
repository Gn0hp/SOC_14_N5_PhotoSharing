package api

import (
	"SOC_N5_14_BTL/internal/repository/flickr_repo"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session/v3"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s Service) AuthorizeFlickr(c *gin.Context) {
	requestToken, requestTokenSecret, err := s.OauthConfig.FlickrConfig.RequestToken()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error request oauthToken: %s", err.Error()))
		return
	}
	authorizationUrl, err := s.OauthConfig.FlickrConfig.AuthorizationURL(requestToken)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting authorizationURL: %s", err.Error()))
		return
	}
	sess, err := session.Start(context.Background(), c.Writer, c.Request)
	sess.Set("flickr_request_token", requestToken)
	sess.Set("flickr_request_token_secret", requestTokenSecret)
	err = sess.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving oauthToken to session: %s", err.Error()))
		return
	}
	c.Redirect(http.StatusFound, authorizationUrl.String())

}
func (s Service) AuthorizeFlickrCallback(c *gin.Context) {

	sess, _ := session.Start(context.Background(), c.Writer, c.Request)
	reqToken, ok := sess.Get("flickr_request_token")
	reqTokenSecret, ok := sess.Get("flickr_request_token_secret")

	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting request token from session: %v", ok))
		return
	}
	verifier := c.Query("oauth_verifier")
	// Exchange code for token
	accessToken, accessTokenSecret, err := s.OauthConfig.FlickrConfig.AccessToken(reqToken.(string), reqTokenSecret.(string), verifier)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error request for access token : %s", err.Error()))
		return
	}
	sess.Set("flickr_access_token", accessToken)
	sess.Set("flickr_access_secret", accessTokenSecret)
	GetUserID(c)
	err = sess.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving oauthToken to session: %s", err.Error()))
		return
	}
	c.String(http.StatusFound, fmt.Sprintf("Success generate token: %v  --- %v", accessToken, accessTokenSecret))
}

func (s Service) FlickrUploadImage(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File["files"]

	flickrRepo := flickr_repo.New(c)
	for _, file := range files {
		openFiles, err := file.Open()
		if err != nil {
			logrus.Errorf("Error openning the file: %v", err)
		}
		defer openFiles.Close()
		buf := make([]byte, file.Size)
		_, err = openFiles.Read(buf)
		if err != nil {
			logrus.Errorf("Error reading the file: %v ", err)
		}
		reader := bytes.NewReader(buf)

		_, ok := flickrRepo.UploadPhoto(reader, name)
		if !ok {
			c.String(http.StatusBadRequest, fmt.Sprint("Error uploading file"))
			return
		}
	}
	logrus.Info(" -------------------", name, email)
	c.String(http.StatusOK, fmt.Sprint("Successfully"))
}

//No need
func (s Service) GetPhotoById(c *gin.Context) {
	id := c.Query("photo_id")
	secret := c.Query("photo_secret")

	repo := flickr_repo.New(c)
	res, err := repo.GetPhotoInfo(id, secret)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res.Photo)
}
func (s Service) GetPhotoByUserId(c *gin.Context) {
	repo := flickr_repo.New(c)
	sess, _ := session.Start(c, c.Writer, c.Request)
	userID, _ := sess.Get("flickr_user_id")
	res, err := repo.GetPhotos(userID.(string))
	if err != nil {
		return
	}

	fmt.Printf("photo: %v", res)

	c.JSON(http.StatusOK, res)
}

func (s Service) CreatePhotoset(c *gin.Context) {
	repo := flickr_repo.New(c)
	title := c.PostForm("title")
	description := c.PostForm("description")
	primaryPhotoId := c.PostForm("primary_photo")

	res, err := repo.CreatePhotoset(title, description, primaryPhotoId)
	if err != nil {
		logrus.Errorf("Error creating photoset: %v", err)
		return
	}
	logrus.Info("response: ", res)
	c.String(http.StatusOK, fmt.Sprint("Successfully create"))

}
func (s Service) AddPhotosToPhotoset(c *gin.Context) {
	repo := flickr_repo.New(c)
	photosetId := c.PostForm("photoset_id")
	photoIds := c.PostFormArray("photo_ids")
	ok, err := repo.AddPhotosToPhotoset(photosetId, photoIds)
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while adding photos to photoset, check log for more infomation: %v", err))
		return
	}
	c.String(http.StatusOK, fmt.Sprint("Successfully"))

}
func (s Service) RemovePhotosFromPhotoset(c *gin.Context) {
	repo := flickr_repo.New(c)
	photosetId := c.PostForm("photoset_id")
	photoIds := c.PostFormArray("photo_ids")
	ok, err := repo.RemovePhotosFromPhotoset(photosetId, photoIds)
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while removing photos from photoset, check log for more infomation: %v", err))
		return
	}
	c.String(http.StatusOK, fmt.Sprint("Successfully"))
}

func (s Service) TestSession(c *gin.Context) {
	sess, _ := session.Start(c, c.Writer, c.Request)
	accessToken, _ := sess.Get("flickr_access_token")
	accessTokenSecret, _ := sess.Get("flickr_access_secret")
	c.String(http.StatusFound, fmt.Sprintf("Success generate token: %v  --- %v", accessToken, accessTokenSecret))
}
