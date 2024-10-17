package db

import (
	"fmt"

	"context"
	"database/sql"

	"ridge/config/abstract"
	_ "ridge/migration/scripts"

	"github.com/go-faster/errors"
	"github.com/pressly/goose/v3"
)

// MigGoose exec goose migrate
func MigGoose(gConfig *abstract.CGlobal) error {
	sqlDb, err := CreateGooseDb(gConfig)
	if err != nil {
		return err
	}

	// base init path
	migConf := gConfig.Migration
	root := gConfig.Root

	// init migrate
	var startId int64
	_ = sqlDb.QueryRow(fmt.Sprintf("SELECT%sFROM %s LIMIT 1", " `id` ", goose.TableName())).Scan(&startId)
	bInitDir := fmt.Sprintf(".%s", migConf.MigrationDir)
	bInitPath := fmt.Sprintf("%s%s", root, migConf.MigrationDir)
	if startId == 0 {
		// add the start version data
		ctx := context.Background()
		_ = goose.UpToContext(ctx, sqlDb, bInitPath, 0, goose.WithNoVersioning())
	}

	// auto migration
	if gConfig.Migration.AutoMigrate {
		// biz migration
		var opts []goose.OptionsFunc
		if migConf.AllowMissing {
			opts = append(opts, goose.WithAllowMissing())
		}

		// up migration
		fmt.Println(fmt.Sprintf("=== goose run directory: %s ......", bInitDir))
		err = goose.Up(sqlDb, bInitPath)
		if err != nil {
			return errors.Errorf("[%s]migration fail: %v", bInitDir, err)
		}
	}

	return nil
}

// CreateGooseDb create goose sql db
func CreateGooseDb(gConfig *abstract.CGlobal) (*sql.DB, error) {
	var err error
	var goSqlDb *sql.DB
	dbConfig := gConfig.Database
	switch dbConfig.Default {
	case "mysql":
		// mysql connection
		timezone := gConfig.App.DefaultTimezone
		mysqlDsn := CreateMysqlDsn(gConfig.Database.Mysql, timezone)
		goSqlDb, err = goose.OpenDBWithDriver("mysql", mysqlDsn)
	default:
		// other connection
		// ......
	}

	return goSqlDb, err
}
