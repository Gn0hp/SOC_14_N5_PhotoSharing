package config

import (
	"golang.org/x/oauth2"
)

type FlickrApiConfig struct {
	APIKey          string
	APISecret       string
	OathToken       string
	OathTokenSecret string
	Endpoint        oauth2.Endpoint
	RedirectURL     string
}
