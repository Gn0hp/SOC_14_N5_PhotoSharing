package authentication

import "github.com/gin-gonic/gin"

func GetAccessToken(c *gin.Context) string {
	if cookies, err := c.Request.Cookie("access_token"); err != nil {
		return cookies.Value
	}
	return ""
}

func ClearAccessToken(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
}
