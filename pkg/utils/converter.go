package utils

import (
	"SOC_N5_14_BTL/internal/entities"
	"encoding/xml"
	"github.com/sirupsen/logrus"
)

func ParsePhotosXmlToPhotoListResponse(xmlString string) *entities.PhotoListResponse {
	var res entities.PhotoListResponse
	err := xml.Unmarshal([]byte(xmlString), &res)
	if err != nil {
		logrus.Errorf("Error while parsing photo to object PhotoListResponse: %v", err)
	}
	return &res
}
