package api

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/internal/repository/flickr_repo"
	"github.com/gin-gonic/gin"
	"gopkg.in/masci/flickr.v2/auth/oauth"
)

func GetUserID(reqToken, reqTokenSecret, accessToken, accessTokenSecret string, c *gin.Context) entities.FlickrUserResponse {
	//sess, _ := session.Start(c, c.Writer, c.Request)
	repo := flickr_repo.NewWithCookie(reqToken, reqTokenSecret, accessToken, accessTokenSecret)
	resp, _ := oauth.CheckToken(repo.Client, repo.Client.OAuthToken)
	return entities.FlickrUserResponse{
		ID:       resp.OAuth.User.ID,
		Username: resp.OAuth.User.Username,
		Fullname: resp.OAuth.User.Fullname,
	}
}
