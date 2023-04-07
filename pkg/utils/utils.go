package utils

import (
	"SOC_N5_14_BTL/internal/constants"
	"encoding/base64"
	"github.com/golang/glog"
	"math/rand"
)

func RandToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		glog.Fatalf("[Gin-OAuth] Failed to read rand: %v", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func RandomString(n int) string {
	letters := []byte(constants.LETTERS_FOR_RANDOM)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
