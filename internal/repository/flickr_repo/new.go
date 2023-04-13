package flickr_repo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session/v3"
	"gopkg.in/masci/flickr.v2"
	"os"
)

func New(c *gin.Context) FlickrRepository {
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

	flickrRepo := FlickrRepository{Client: client}
	return flickrRepo
}
