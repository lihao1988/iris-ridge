package file

import (
	"os"
	"path/filepath"
)

// GetPathFiles traverse files in the directory
func GetPathFiles(fPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(fPath, func(fName string, info os.FileInfo, err error) error {
		// if not file, exec return
		if info == nil || info.IsDir() {
			return nil
		}

		// append file
		files = append(files, fName)

		return nil
	})

	return files, err
}
