package config

import (
	"SOC_N5_14_BTL/internal/services/databases/mysql"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

type OathConfig struct {
	GoogleOauthConfig *oauth2.Config
	FlickrConfig      *oauth1.Config
}

type Configuration struct {
	Server   ServerConfiguration `mapstructure:"server"`
	Database mysql.MySqlConfig   `mapstructure:"database"`
}

var config *Configuration
var ClientOauthConfig OathConfig

func SetupConfiguration() {
	var configuration *Configuration

	if err := viper.Unmarshal(&configuration); err != nil {
		logrus.Fatalf("Unable to decode into struct: %v", err)
	}
	config = configuration
}
func GetConfig() *Configuration {
	return config
}
func InitConfig() OathConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return OathConfig{
		GoogleOauthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes:       []string{"profile", "email"},
		},
		FlickrConfig: &oauth1.Config{
			ConsumerKey:    os.Getenv("FLICKR_API_KEY"),
			ConsumerSecret: os.Getenv("FLICKR_API_SECRET"),
			CallbackURL:    os.Getenv("FLICKR_REDIRECT_URL"),
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: os.Getenv("FLICKR_REQUEST_TOKEN_URL"),
				AuthorizeURL:    os.Getenv("FLICKR_AUTH_URL"),
				AccessTokenURL:  os.Getenv("FLICKR_TOKEN_URL"),
			},
		},
	}
}
