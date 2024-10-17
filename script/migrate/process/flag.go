package process

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"database/sql"
	"runtime/debug"

	"ridge/config/abstract"

	"github.com/go-faster/errors"
	"github.com/lihao1988/php2go/file"
	"github.com/pressly/goose/v3"
)

var version string

var (
	flags        = flag.NewFlagSet("goose", flag.ExitOnError)
	dir          = flags.String("dir", ".", "directory with migration files")
	table        = flags.String("table", "goose_db_version", "migrations table name")
	verbose      = flags.Bool("v", false, "enable verbose mode")
	help         = flags.Bool("h", false, "print help")
	versionFlag  = flags.Bool("version", false, "print version")
	sequential   = flags.Bool("s", false, "use sequential numbering for new migrations")
	allowMissing = flags.Bool("allow-missing", false, "applies missing (out-of-order) migrations")
	noVersioning = flags.Bool("no-versioning", false, "apply migration commands with no versioning, in file order, from directory pointed to")
	noColor      = flags.Bool("no-color", false, "disable color output (NO_COLOR env variable supported)")
	timeout      = flags.Duration("timeout", 0, "maximum allowed duration for queries to run; e.g., 1h13m")
)

// gConfig global config
var gConfig *abstract.CGlobal

// FlagRun run flag
func FlagRun(ctx context.Context, gcf *abstract.CGlobal, db *sql.DB) ([]string, error) {
	// set global config
	SetGConfig(gcf)

	// flag operate
	args, err := flagOpt()
	if err != nil {
		return nil, err
	}
	if len(args) == 0 {
		flags.Usage()
		os.Exit(1)
	}

	// switch operate
	err = SwitchOpt(ctx, db, args)
	if err != nil {
		return nil, err
	}

	return args, nil
}

// flagOpt flag operate
func flagOpt() ([]string, error) {
	// flag usage
	flags.Usage = usage
	if err := flags.Parse(os.Args[1:]); err != nil {
		return nil, errors.Errorf("failed to parse args: %v", err)
	}

	// flag args
	// goose
	if *versionFlag {
		buildInfo, ok := debug.ReadBuildInfo()
		if version == "" && ok && buildInfo != nil && buildInfo.Main.Version != "" {
			version = buildInfo.Main.Version
		}
		fmt.Printf("goose version: %s\n", strings.TrimSpace(version))
		return nil, nil
	}
	if *verbose {
		goose.SetVerbose(true)
	}
	if *sequential {
		goose.SetSequential(true)
	}
	goose.SetTableName(*table)

	// flag args
	args := flags.Args()

	if *help {
		flags.Usage()
		return nil, nil
	}

	return args, nil
}

// SwitchOpt switch operate
func SwitchOpt(ctx context.Context, db *sql.DB, args []string) error {
	// check migration directory
	dirPath := getDirPath()
	if !file.FileExists(dirPath) {
		return errors.New("goose run dir is not exist !")
	}

	// switch operate
	var err error
	switch args[0] {
	case "init":
		if err = gooseInit(dirPath); err != nil {
			err = errors.Errorf("goose run: %v", err)
		}
		return err
	case "create":
		if err = goose.RunContext(ctx, "create", nil, dirPath, args[1:]...); err != nil {
			err = errors.Errorf("goose run: %v", err)
		}
		return err
	case "fix":
		if err = goose.RunContext(ctx, "fix", nil, dirPath); err != nil {
			err = errors.Errorf("goose run: %v", err)
		}
		return err
	case "validate":
		if err = printValidate(dirPath, *verbose); err != nil {
			err = errors.Errorf("goose validate: %v", err)
		}
		return err
	case "status":
		if err = goose.Status(db, dirPath); err != nil {
			err = errors.Errorf("goose status: %v", err)
		} else {
			os.Exit(1)
		}

		return err
	}

	if len(args) == 0 {
		flags.Usage()
		os.Exit(1)
	}

	return nil
}

// CmdRun run command
func CmdRun(ctx context.Context, db *sql.DB, args []string) error {
	command := args[0]

	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	var options []goose.OptionsFunc
	if *noColor || checkNoColorFromEnv() {
		options = append(options, goose.WithNoColor(true))
	}
	if *allowMissing {
		options = append(options, goose.WithAllowMissing())
	}
	if *noVersioning {
		options = append(options, goose.WithNoVersioning())
	}
	if timeout != nil && *timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	var err error
	if err = goose.RunWithOptionsContext(
		ctx,
		command,
		db,
		getDirPath(),
		arguments,
		options...,
	); err != nil {
		err = errors.Errorf("goose run: %v", err)
	}

	return err
}

// SetGConfig set global config
func SetGConfig(gcf *abstract.CGlobal) {
	gConfig = gcf
}

// usage show operate context
func usage() {
	fmt.Println(usagePrefix, "")
	flags.PrintDefaults()
	fmt.Println(usageCommands, "")
}

// getDirPath get dir path
func getDirPath() string {
	dirPath := *dir
	if dirPath == "." {
		dirPath = gConfig.Root + gConfig.Migration.MigrationDir
	}
	if strings.Contains(dirPath, "./") {
		dirPath = gConfig.Root + strings.Replace(dirPath, "./", "/", 1)
	}

	return dirPath
}
