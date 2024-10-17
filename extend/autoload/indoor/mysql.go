package indoor

import (
	"gorm.io/gorm"
	"ridge/common/global"
	"ridge/common/lib/db"
	"ridge/config/abstract"
)

// AutoloadDb autoload db connection
func AutoloadDb() error {
	// set global database
	gConfig := global.GConfig
	switch gConfig.Database.Default {
	case "mysql":
		gormDb, err := LoadMysqlDb(gConfig)
		if err != nil {
			return err
		}
		global.GDB = gormDb
	default:
		// other connection
		// ......
	}

	return nil

}

// LoadMysqlDb load database
func LoadMysqlDb(gConfig *abstract.CGlobal) (*gorm.DB, error) {
	// mysql connection
	timezone := gConfig.App.DefaultTimezone
	gormDb, err := db.InitGormMysql(gConfig.Database.Mysql, timezone)
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}
