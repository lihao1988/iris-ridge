package process

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"

	"path/filepath"
	"text/tabwriter"

	"ridge/script/migrate/process/migrationstats"

	"github.com/pressly/goose/v3"
)

// initDir will create a directory with an empty SQL migration file.
func gooseInit(dir string) error {
	if dir == "" || dir == "." {
		dir = "migrations"
	}
	_, err := os.Stat(dir)
	switch {
	case errors.Is(err, fs.ErrNotExist):
	case err == nil, errors.Is(err, fs.ErrExist):
		return fmt.Errorf("directory already exists: %s", dir)
	default:
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return goose.CreateWithTemplate(nil, dir, sqlMigrationTemplate, "initial", "sql")
}

func gatherFilenames(filename string) ([]string, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	var filenames []string
	if stat.IsDir() {
		for _, pattern := range []string{"*.sql", "*.go"} {
			file, err := filepath.Glob(filepath.Join(filename, pattern))
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, file...)
		}
	} else {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)
	return filenames, nil
}

func printValidate(filename string, verbose bool) error {
	filenames, err := gatherFilenames(filename)
	if err != nil {
		return err
	}
	stats, err := migrationstats.GatherStats(
		migrationstats.NewFileWalker(filenames...),
		false,
	)
	if err != nil {
		return err
	}
	// TODO(mf): we should introduce a --debug flag, which allows printing
	// more internal debug information and leave verbose for additional information.
	if !verbose {
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmtPattern := "%v\t%v\t%v\t%v\t%v\t\n"
	_, _ = fmt.Fprintf(w, fmtPattern, "Type", "Txn", "Up", "Down", "Name")
	_, _ = fmt.Fprintf(w, fmtPattern, "────", "───", "──", "────", "────")
	for _, m := range stats {
		txnStr := "✔"
		if !m.Tx {
			txnStr = "✘"
		}
		_, _ = fmt.Fprintf(w, fmtPattern,
			strings.TrimPrefix(filepath.Ext(m.FileName), "."),
			txnStr,
			m.UpCount,
			m.DownCount,
			filepath.Base(m.FileName),
		)
	}
	return w.Flush()
}
