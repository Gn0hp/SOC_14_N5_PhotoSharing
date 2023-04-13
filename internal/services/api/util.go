package api

import (
	"SOC_N5_14_BTL/internal/repository/flickr_repo"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session/v3"
	"gopkg.in/masci/flickr.v2/auth/oauth"
)

func GetUserID(c *gin.Context) {
	sess, _ := session.Start(c, c.Writer, c.Request)
	repo := flickr_repo.New(c)
	resp, _ := oauth.CheckToken(repo.Client, repo.Client.OAuthToken)
	sess.Set("flickr_user_id", resp.OAuth.User.ID)
}
