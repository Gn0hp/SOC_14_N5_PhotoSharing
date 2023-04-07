package server

import (
	"SOC_N5_14_BTL/cmd/server/routes"
	"SOC_N5_14_BTL/internal/log"
	"SOC_N5_14_BTL/internal/services"
	"SOC_N5_14_BTL/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

//ctx := context.TODO()
//v := viper.New()
func Start() {
	var defaultSrv services.DefaultService
	defaultSrv.Init()

	var logCf log.LogConfig
	cfByte, _ := json.Marshal(viper.GetStringMap("log"))
	json.Unmarshal(cfByte, &logCf)
	logger := log.NewLogger(logCf)
	log.SetStandaloneLogger(logger)

	config.SetupConfiguration()

	if err := run(); err != nil {
		logrus.Fatal(err)
	}

}

func run() error {
	env := os.Getenv("EXECUTE_ENV")
	port := 8900
	var host string
	if env == "dev" {
		host = os.Getenv("DEV_HOST_API")
	} else if env == "docker" {
		host = os.Getenv("DOCKER_HOST_API")
	} else if env == "prod" {
		host = os.Getenv("PROD_HOST_API")
	}
	r := routes.Setup()
	//err := r.Run("127.0.0.1:8900")
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logrus.Panic(err)
		return err
	}
	return nil
}
