package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnector(config MySqlConfig) (*gorm.DB, error) {
	config.Params = make(map[string]string)
	config.Params["parseTime"] = "true"
	config.Params["rejectReadOnly"] = "true"
	db, err := gorm.Open(mysql.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connect database failed, error: %v", err))
	}
	return db, nil
}
