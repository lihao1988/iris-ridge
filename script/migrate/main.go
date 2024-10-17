package main

import (
	"context"
	_ "embed"
	"log"

	config "ridge/config/abstract"
	_ "ridge/migration/scripts"
	"ridge/script/migrate/process"
	"ridge/tool"
)

// migrate
func main() {
	// context background
	ctx := context.Background()

	// load config
	var err error
	gConfig := &config.CGlobal{}
	gConfig.Root = tool.GetExecFilePath()
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
	defer func() {
		if dbErr := sqlDb.Close(); dbErr != nil {
			log.Fatalf("goose: failed to close DB: %v\n", dbErr)
		}
	}()

	// flag operate
	args, err := process.FlagRun(ctx, gConfig, sqlDb)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// run cmd
	err = process.CmdRun(ctx, sqlDb, args)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
