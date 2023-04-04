package api

import "SOC_N5_14_BTL/pkg/config"

type Service struct {
	GoogleOauthConfig *config.OathConfig
}

func NewService(oc config.OathConfig) Service {
	return Service{
		GoogleOauthConfig: &oc,
	}
}
