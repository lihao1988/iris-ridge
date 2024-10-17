package main

import (
	"context"
	"log"
	"testing"

	config "ridge/config/abstract"
	_ "ridge/migration/scripts"
	"ridge/script/migrate/process"

	"github.com/pressly/goose/v3"
)

func TestMigrate(t *testing.T) {
	// load config
	var err error
	gConfig := &config.CGlobal{}
	gConfig.Root = "/home/lgmin1988/go_project/ridge"
	err = process.LoadConfig(gConfig)
	if err != nil {
		log.Fatalf("load config: %v", err)
		return
	}

	// load sql db
	sqlDb, err := process.LoadSqlDb(gConfig)
	if err != nil {
		log.Fatalf("load sqlDb: %v", err)
		return
	}

	// migration status
	err = goose.Status(sqlDb, "/home/lgmin1988/go_project/ridge/migration/scripts")
	if err != nil {
		return
	}
}

func TestValidate(t *testing.T) {
	// load config
	var err error
	gConfig := &config.CGlobal{}
	gConfig.Root = "/home/lgmin1988/go_project/ridge"
	err = process.LoadConfig(gConfig)
	if err != nil {
		log.Fatalf("load config: %v", err)
		return
	}

	// load sql db
	sqlDb, err := process.LoadSqlDb(gConfig)
	if err != nil {
		log.Fatalf("load sqlDb: %v", err)
		return
	}

	// flag operate
	args := []string{"validate"}
	ctx := context.Background()
	process.SetGConfig(gConfig)
	err = process.SwitchOpt(ctx, sqlDb, args)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
