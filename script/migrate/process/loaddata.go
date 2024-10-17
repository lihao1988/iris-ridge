package process

import (
	"database/sql"
	"ridge/common/lib/db"
	"ridge/config/abstract"
	"ridge/extend/autoload/indoor"
)

// LoadConfig load config data
func LoadConfig(gConfig *abstract.CGlobal) error {
	return indoor.LoadConfig(gConfig)
}

// LoadSqlDb create sql.db
func LoadSqlDb(gConfig *abstract.CGlobal) (*sql.DB, error) {
	return db.CreateGooseDb(gConfig)
}
