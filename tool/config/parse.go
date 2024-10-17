package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"encoding/json"
	"github.com/lihao1988/php2go/array"
	"github.com/lihao1988/php2go/file"
	"gopkg.in/yaml.v3"
)

// SupportedExtList are universally supported extensions.
var SupportedExtList = []string{"json", "yaml", "yml"}

// Parse parse config data
func Parse(fPath string, out interface{}) error {
	// get file data
	dataBytes, fExt, err := readConfig(fPath)
	if err != nil {
		return err
	}

	outType := reflect.TypeOf(out)
	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr || outValue.IsNil() {
		return errors.New(fmt.Sprintf("haystack: out type[%s] error", outType))
	}
	v := outValue.Convert(outType).Interface()

	// switch file type
	switch strings.ToLower(fExt) {
	case "yaml", "yml":
		if err := yaml.Unmarshal(dataBytes, v); err != nil {
			return ConfigParseError{err}
		}

	case "json":
		if err := json.Unmarshal(dataBytes, v); err != nil {
			return ConfigParseError{err}
		}

	}
	out = v

	return nil
}

// readConfig read config file
func readConfig(fPath string) ([]byte, string, error) {
	isExist := file.FileExists(fPath)
	if !isExist {
		return nil, "", errors.New("file is not exist!")
	}

	dataBytes, err := os.ReadFile(fPath)
	if err != nil {
		return nil, "", err
	}

	ext := path.Ext(fPath)
	if len(ext) > 1 {
		ext = ext[1:]
	}
	if !array.In(ext, SupportedExtList) {
		return nil, "", errors.New("file is not support!")
	}

	return dataBytes, ext, nil
}
