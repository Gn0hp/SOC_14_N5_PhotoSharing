package server

import (
	"SOC_N5_14_BTL/cmd/server/routes"
	"github.com/sirupsen/logrus"
)

//ctx := context.TODO()
//v := viper.New()
func Start() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	r := routes.Setup()
	err := r.Run("127.0.0.1:8900")
	if err != nil {
		logrus.Panic(err)
		return err
	}

	return nil
}
