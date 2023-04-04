package routes

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/internal/services/api"
	"SOC_N5_14_BTL/pkg/config"
	"SOC_N5_14_BTL/pkg/middlewares"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
)

func initConfig() config.OathConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return config.OathConfig{
		GoogleOauthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			RedirectURL:  "http://localhost:8900/auth/google/callback",
			Scopes:       []string{"profile", "email"},
		},

		FlickrConfig: config.FlickrApiConfig{
			APIKey:          os.Getenv("FLICKR_API_KEY"),
			APISecret:       os.Getenv("FLICKR_API_SECRET"),
			OathToken:       os.Getenv("FLICKR_OAUTH_TOKEN"),
			OathTokenSecret: os.Getenv("FLICKR_OAUTH_TOKEN_SECRET"),
		},
	}
}
func Setup() *gin.Engine {
	cfg := initConfig()
	r := gin.New()
	srv := api.NewService(cfg)
	r.Use(middlewares.JSONMiddleware())
	r.Use(middlewares.CORS())

	googleAuth := r.Group("/auth/google")
	{
		//redirect to sign in with google page
		googleAuth.GET("", srv.DirectToAuthenPage)

		// after sign in, google redirect to this with user info
		googleAuth.GET("/callback", func(c *gin.Context) {
			code := c.Query("code")

			// Exchange code for token
			token, err := cfg.GoogleOauthConfig.Exchange(context.Background(), code)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Error exchanging code: %s", err.Error()))
				return
			}
			client := cfg.GoogleOauthConfig.Client(c.Request.Context(), token)

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

			logrus.Infof("User info received: %s with id %s", userInfo.Email, userInfo.ID)

		})
		googleAuth.GET("/logout", func(c *gin.Context) {

		})
	}
	return r
}
