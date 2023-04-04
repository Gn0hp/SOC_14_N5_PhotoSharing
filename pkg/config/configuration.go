package config

import "golang.org/x/oauth2"

type OathConfig struct {
	GoogleOauthConfig *oauth2.Config
	FlickrConfig      FlickrApiConfig
}
