package interfaces

import "github.com/gin-gonic/gin"

type ServiceInterface interface {
	AuthorizeFlickr(c *gin.Context)
	AuthorizeFlickrCallback(c *gin.Context)

	FlickrUploadImage(c *gin.Context)
	GetPhotoById(c *gin.Context)
	GetPhotoByUserId(c *gin.Context)

	CreatePhotoset(c *gin.Context)
	AddPhotosToPhotoset(c *gin.Context)
	RemovePhotosFromPhotoset(c *gin.Context)
}
