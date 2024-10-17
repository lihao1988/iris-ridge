package db

import (
	"errors"
	"fmt"
	"strings"

	"ridge/config/abstract/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/lihao1988/php2go/url"
)

// InitGormMysql init gorm mysql connection
func InitGormMysql(dbConfig *structs.MysqlConn, timezone string) (*gorm.DB, error) {
	dsn := CreateMysqlDsn(dbConfig, timezone)
	gormConfig := &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: dbConfig.Prefix, // 设置表前缀
	}}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, errors.New("mysql connection fail ~")
	}

	// set param
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Mysql connection error:%s", err.Error()))
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleCount)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenCount)

	return db, nil
}

// CreateMysqlDsn create mysql dsn
func CreateMysqlDsn(dbConfig *structs.MysqlConn, timezone string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
		dbConfig.Charset,
		url.UrlEncode(timezone),
	)

	// extend params
	if len(dbConfig.Params) > 1 {
		var params []string
		for key, val := range dbConfig.Params {
			params = append(params, fmt.Sprintf("%s=%s", key, url.UrlEncode(val)))
		}
		dsn = dsn + "&" + strings.Join(params, "&")
	}

	return dsn
}
