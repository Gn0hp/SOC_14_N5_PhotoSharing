package cli

import (
	"SOC_N5_14_BTL/internal/entities"
	"SOC_N5_14_BTL/internal/services"
	"SOC_N5_14_BTL/internal/services/databases/mysql"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type MigrateService struct {
	services.DefaultService
	Context context.Context
	GormDB  *gorm.DB
}

func Migrate() {
	migrateService := MigrateService{}
	migrateService.Init()

	tables := []interface{}{
		entities.PhotoIdUrlMapping{},
	}
	err := migrateService.GormDB.AutoMigrate(tables...)
	if err != nil {
		logrus.Errorf("Error migrate to database: %v", err)
		return
	}
	logrus.Info("Database seed successfully")
}

func (m *MigrateService) Init() {
	m.DefaultService.Init()
	dbCf := mysql.MySqlConfig{}

	cfBytes, _ := json.Marshal(viper.GetStringMap("database"))
	json.Unmarshal(cfBytes, &dbCf)

	gormDb, err := mysql.NewConnector(dbCf)
	if err != nil {
		panic(err)
	}
	m.GormDB = gormDb
}
