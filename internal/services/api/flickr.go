package api

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/internal/repository/flickr_repo"
	"bytes"
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
	//@dev NOTE: Use for PC, not android
	//sess, err := session.Start(c, c.Writer, c.Request)
	//sess.Set("flickr_request_token", requestToken)
	//sess.Set("flickr_request_token_secret", requestTokenSecret)
	//err = sess.Save()

	//@dev NOTE: Use for android
	c.SetCookie("gin_cookie_frt", requestToken, 360, "/", "10.0.2.2:8900", false, true)
	c.SetCookie("gin_cookie_frts", requestTokenSecret, 360, "/", "10.0.2.2:8900", false, true)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving oauthToken to session: %s", err.Error()))
		return
	}
	fmt.Printf("Go session ID first : %v", c.Request.Header.Get("Cookie"))
	c.Redirect(http.StatusFound, authorizationUrl.String())
}
func (s Service) AuthorizeFlickrCallback(c *gin.Context) {
	//@dev NOTE: Use for PC, not android
	sess, _ := session.Start(c, c.Writer, c.Request)
	//reqToken, ok := sess.Get("flickr_request_token")
	//reqTokenSecret, ok := sess.Get("flickr_request_token_secret")
	//fmt.Printf("Go session ID second : %v", c.Request.Header.Get("Cookie"))
	reqToken, err := c.Cookie("gin_cookie_frt")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting request token from cookie: %v", err))
		return
	}
	reqTokenSecret, err := c.Cookie("gin_cookie_frts")

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting request token secret from cookie: %v", err))
		return
	}
	verifier := c.Query("oauth_verifier")
	// Exchange code for token
	accessToken, accessTokenSecret, err := s.OauthConfig.FlickrConfig.AccessToken(reqToken, reqTokenSecret, verifier)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error request for access token : %s", err.Error()))
		return
	}
	sess.Set("flickr_request_token", reqToken)
	sess.Set("flickr_request_token_secret", reqTokenSecret)
	sess.Set("flickr_access_token", accessToken)
	sess.Set("flickr_access_secret", accessTokenSecret)
	user := GetUserID(reqToken, reqTokenSecret, accessToken, accessTokenSecret, c)

	//sess.Set("flickr_user_id", userId)
	c.SetCookie("flickr_user_id", user.ID, 360, "/", "10.0.2.2:8900", false, false)
	c.SetCookie("flickr_user_username", user.Username, 360, "/", "10.0.2.2:8900", false, false)
	c.SetCookie("flickr_user_fullname", user.Fullname, 360, "/", "10.0.2.2:8900", false, false)
	c.SetCookie("flickr_user_id", user.ID, 360, "/", "localhost:8900", false, false)
	c.SetCookie("flickr_user_username", user.Username, 360, "/", "localhost:8900", false, false)
	c.SetCookie("flickr_user_fullname", user.Fullname, 360, "/", "localhost:8900", false, false)
	err = sess.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving oauthToken to session: %s", err.Error()))
		return
	}
	//c.JSON(http.StatusOK, map[string]string{
	//	"user_id":  user.ID,
	//	"username": user.Username,
	//	"fullname": user.Fullname,
	//})
	c.Redirect(http.StatusFound, fmt.Sprintf("http://localhost:8900/flickr/redirect-callback?"+
		"user_id=%s"+
		"&username=%s"+
		"&fullname=%s"+
		"&flickr_request_token=%s"+
		"&flickr_request_token_secret=%s"+
		"&flickr_access_token=%s"+
		"&flickr_access_secret=%s", user.ID, user.Username, user.Fullname, reqToken, reqTokenSecret, accessToken, accessTokenSecret))

}

func (s Service) FlickrUploadImage(c *gin.Context) {
	fmt.Println("go here ...")
	username, _ := c.Cookie("flickr_user_username")
	id, _ := c.Cookie("flickr_user_id")
	reqToken, _ := c.Cookie("flickr_request_token")
	reqTokenSecret, _ := c.Cookie("flickr_request_token_secret")
	accessToken, _ := c.Cookie("flickr_access_token")
	accessTokenSecret, _ := c.Cookie("flickr_access_secret")

	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File["files"]

	var uploadedPhotoId []string

	flickrRepo := flickr_repo.NewWithCookie(reqToken, reqTokenSecret, accessToken, accessTokenSecret)
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

		response, ok := flickrRepo.UploadPhoto(reader, username)
		if !ok {
			c.String(http.StatusBadRequest, fmt.Sprint("Error uploading file"))
			return
		}
		uploadedPhotoId = append(uploadedPhotoId, response.ID)
	}
	resp, err := flickrRepo.GetPhotos(id)
	var res []entities.PhotoResponse
	for _, photo := range resp.Photo {
		if contain(uploadedPhotoId, photo.Id) {
			res = append(res, photo)
		}
	}
	c.JSON(http.StatusOK, res)
}
func contain(arr []string, element string) bool {
	for _, val := range arr {
		if val == element {
			return true
		}
	}
	return false
}
func (s Service) UploadPost(c *gin.Context) {

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
