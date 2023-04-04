package mysql

import (
	"fmt"
	"gorm.io/gorm"
)

type DB struct {
	gormDB *gorm.DB
}

func New(config MySqlConfig) *DB {
	db, err := NewConnector(config)
	if err != nil {
		panic(fmt.Sprintf("connect database failed, error: %v", err))
	}
	return &DB{
		gormDB: db,
	}
}
func (db *DB) GormDB() *gorm.DB {
	return db.gormDB
}

func (db *DB) Close() {
	dbConnection, _ := db.gormDB.DB()
	if dbConnection != nil {
		dbConnection.Close()
	}

}
