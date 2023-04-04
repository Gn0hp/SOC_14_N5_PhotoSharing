package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang/glog"
)

func RandToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		glog.Fatalf("[Gin-OAuth] Failed to read rand: %v", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
