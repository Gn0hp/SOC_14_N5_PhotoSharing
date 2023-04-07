package api

import (
	"SOC_N5_14_BTL/internal/repository"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/masci/flickr.v2"
	"net/http"
	"os"
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

	sess, _ := session.Start(context.Background(), c.Writer, c.Request)

	reqToken, _ := sess.Get("flickr_request_token")
	reqTokenSecret, _ := sess.Get("flickr_request_token_secret")
	accessToken, _ := sess.Get("flickr_access_token")
	accessTokenSecret, _ := sess.Get("flickr_access_secret")

	client := flickr.NewFlickrClient(reqToken.(string), reqTokenSecret.(string))
	client.OAuthToken = accessToken.(string)
	client.OAuthTokenSecret = accessTokenSecret.(string)
	client.ApiKey = os.Getenv("FLICKR_API_KEY")
	client.ApiSecret = os.Getenv("FLICKR_API_SECRET")

	flickrRepo := repository.FlickrRepository{Client: client}
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

func (s Service) TestSession(c *gin.Context) {
	sess, _ := session.Start(c, c.Writer, c.Request)
	accessToken, _ := sess.Get("flickr_access_token")
	accessTokenSecret, _ := sess.Get("flickr_access_secret")
	c.String(http.StatusFound, fmt.Sprintf("Success generate token: %v  --- %v", accessToken, accessTokenSecret))
}
