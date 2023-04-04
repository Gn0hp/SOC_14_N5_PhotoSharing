package services

import (
	"github.com/sirupsen/logrus"
)

var AppManagement map[string]interface{}

func init() {
	if AppManagement == nil {
		AppManagement = make(map[string]interface{})
		logrus.Info("Init new app management")
	}
}
