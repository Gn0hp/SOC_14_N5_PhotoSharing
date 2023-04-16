package routes

import (
	"SOC_N5_14_BTL/internal/services/api"
	"SOC_N5_14_BTL/pkg/config"
	"SOC_N5_14_BTL/pkg/middlewares"
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func Setup() *gin.Engine {
	config.ClientOauthConfig = config.InitConfig()
	certPEMBlock, keyPemBlock := ReadPemFile()
	cert := tls.Certificate{
		Certificate:                  [][]byte{certPEMBlock.Bytes},
		PrivateKey:                   keyPemBlock.Bytes,
		SupportedSignatureAlgorithms: nil,
		OCSPStaple:                   nil,
		SignedCertificateTimestamps:  nil,
		Leaf:                         nil,
	}

	r := gin.New()
	srv := api.NewService(config.ClientOauthConfig)
	r.Use(middlewares.JSONMiddleware())
	r.Use(middlewares.CORS())

	rSecure := &http.Server{
		Addr:    ":8901",
		Handler: r,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	go func() {
		err := rSecure.ListenAndServeTLS("localhost.pem", "localhost-key.pem")
		if err != nil {
			log.Fatalf("Error while listen and server https server with tls: %v", err)
		}
	}()
	//r.Static("/public/static/", "./public/static")
	r.LoadHTMLGlob("public/*")

	apiEp := r.Group("/api/v1")
	{
		photoEp := apiEp.Group("photo")
		{
			photoEp.GET("/getById", srv.GetPhotoById) // -> c.Param c.Param
			photoEp.GET("/getByUserId", srv.GetPhotoByUserId)
		}
		photosetEp := apiEp.Group("/photoset")
		{
			photosetEp.GET("/create", func(c *gin.Context) {
				c.Header("Content-Type", "text/html")
				c.HTML(http.StatusOK, "createPhotoset.html", gin.H{
					"ActionPath": "/api/v1/photoset/create",
				})
			})
			photosetEp.POST("/create", srv.CreatePhotoset)
			photosetEp.POST("/addToPhotoset", srv.AddPhotosToPhotoset)
			photosetEp.POST("removeFromPhotoset", srv.RemovePhotosFromPhotoset)

		}
	}
	googleAuth := r.Group("/auth/google")
	{
		//redirect to sign in with google page
		googleAuth.GET("", srv.DirectToGoogleAuthenPage)
		// after sign in, google redirect to this with user info
		googleAuth.GET("/callback", srv.GoogleSigninCallback)
		googleAuth.GET("/logout", srv.GoogleLogout)
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusAccepted, fmt.Sprint("Welcome"))
	})
	flickrAuth := r.Group("/flickr")
	{
		flickrAuth.GET("/callbackAuth", srv.AuthorizeFlickrCallback)
		flickrAuth.GET("/auth", srv.AuthorizeFlickr)
		flickrAuth.GET("/testSession", srv.TestSession)
		flickrAuth.GET("/upload", func(c *gin.Context) {
			c.Header("Content-Type", "text/html")
			c.HTML(http.StatusOK, "uploadFile.html", gin.H{
				"ActionPath": "/flickr/upload_img",
			})
		})
		flickrAuth.POST("/upload_img", srv.FlickrUploadImage)
	}

	return r
}

func ReadPemFile() (*pem.Block, *pem.Block) {
	certPEM, err := ioutil.ReadFile("localhost.pem")
	if err != nil {
		log.Fatal(err)
	}
	keyPEM, err := ioutil.ReadFile("localhost-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPEMBlock, _ := pem.Decode(certPEM)
	keyPEMBlock, _ := pem.Decode(keyPEM)
	return certPEMBlock, keyPEMBlock
}
