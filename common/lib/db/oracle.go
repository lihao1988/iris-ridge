package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/godoes/gorm-oracle"
)

// InitGormOracle init gorm oracle connection
func InitGormOracle() (*gorm.DB, error) {
	dsn := CreateOracleDsn()
	db, err := gorm.Open(oracle.New(oracle.Config{
		DSN:                 dsn,
		IgnoreCase:          true,
		NamingCaseSensitive: true,
	}), getGormConfig())
	if err != nil {
		return nil, errors.New("oracle connection fail ~")
	}

	// set param
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Mysql connection error:%s", err.Error()))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil

}

// CreateOracleDsn create Oracle dsn
func CreateOracleDsn() string {
	language := "SIMPLIFIED CHINESE"
	territory := "CHINA"
	dsn := oracle.BuildUrl("172.168.2.62", 1521,
		"ora_csdn",
		"system",
		"system",
		map[string]string{
			"CONNECTION TIMEOUT": "90",
			"LANGUAGE":           language,
			"TERRITORY":          territory,
			"SSL":                "false",
		},
	)

	return dsn
}

// getGormConfig set gorm logger config
func getGormConfig() *gorm.Config {
	logWriter := new(log.Logger)
	logWriter.SetOutput(os.Stdout)

	return &gorm.Config{
		Logger: logger.New(
			logWriter,
			logger.Config{LogLevel: logger.Info},
		),
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
	}
}
