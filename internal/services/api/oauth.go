package api

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/pkg/config"
	"SOC_N5_14_BTL/pkg/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session/v3"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (s Service) DirectToGoogleAuthenPage(c *gin.Context) {
	url := s.OauthConfig.GoogleOauthConfig.AuthCodeURL(utils.RandomString(32))
	c.Redirect(http.StatusFound, url)
}

func (s Service) GoogleSigninCallback(c *gin.Context) {
	code := c.Query("code")

	// Exchange code for token
	token, err := config.ClientOauthConfig.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error exchanging code: %s", err.Error()))
		return
	}
	client := config.ClientOauthConfig.GoogleOauthConfig.Client(c.Request.Context(), token)

	// Get user info
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting user info: %s", err.Error()))
		return
	}

	defer resp.Body.Close()
	userInfo := entities.ResponseGoogleUserInfo{}

	var byt bytes.Buffer
	_, _ = io.Copy(&byt, resp.Body)
	if err := json.Unmarshal(byt.Bytes(), &userInfo); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error decoding user info: %s", err.Error()))
		return
	}

	logrus.Infof("User info received: %v ", userInfo)
	sess, err := session.Start(context.Background(), c.Writer, c.Request)
	sess.Set("token_user", userInfo.ID)
	err = sess.Save()
	if err != nil {
		return
	}
}
func (s Service) GoogleLogout(c *gin.Context) {

}
