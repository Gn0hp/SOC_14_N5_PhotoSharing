package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Service) DirectToAuthenPage(c *gin.Context) {
	url := s.GoogleOauthConfig.GoogleOauthConfig.AuthCodeURL("state-token")
	c.Redirect(http.StatusFound, url)
}
