package indoor

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"ridge/common/global"
	"ridge/config/abstract"
	"ridge/config/abstract/structs"
	tConfig "ridge/tool/config"
	tFile "ridge/tool/file"

	"github.com/lihao1988/php2go/file"
	"github.com/spf13/cast"
)

// config const data
const (
	// configDir config dir
	configDir = "/config"

	// appConfigFile app config file
	appConfigFile = "app.yml"
)

// AutoloadConfig parsing the path of config
func AutoloadConfig() error {
	return LoadConfig(global.GConfig)
}

// LoadConfig load config data
func LoadConfig(gConfig *abstract.CGlobal) error {
	configPath := fmt.Sprintf("%s%s", gConfig.Root, configDir)

	// app.yml (system config, first parse)
	appData := &structs.App{}
	appFileName := fmt.Sprintf("%s/%s", configPath, appConfigFile)
	err := tConfig.Parse(appFileName, appData)
	gConfig.App = appData

	// traverse the path of config
	evnConfigPath := fmt.Sprintf("%s/%s", configPath, appData.Env)
	files, err := tFile.GetPathFiles(evnConfigPath)
	for _, fName := range files {
		err = parseForFile(fName, gConfig)
		if err != nil {
			break
		}
	}

	return err
}

// parseForFile parse the config file
func parseForFile(fName string, gConfig *abstract.CGlobal) error {
	fInfo, err := file.PathInfo(fName)
	if err != nil {
		return nil
	}
	pFileName := strings.Replace(cast.ToString(fInfo["filename"]),
		cast.ToString(fInfo["extension"]), "", 1)
	cInstance, ok := structs.RegConfigMap[pFileName]
	instAttr, _ := abstract.CGlobalAttrMap[pFileName]
	if !ok {
		return errors.New(fmt.Sprintf("config[%s] instance is not exist!", pFileName))
	}

	instType := reflect.TypeOf(cInstance)
	instValue := reflect.ValueOf(cInstance)
	if instValue.Kind() != reflect.Ptr || instValue.IsNil() {
		return errors.New(fmt.Sprintf("haystack: out type[%s] error", instType))
	}
	inst := instValue.Convert(instType).Interface()
	err = tConfig.Parse(fName, inst)
	if err != nil {
		return err
	}

	gField := reflect.ValueOf(gConfig).Elem().FieldByName(instAttr)
	gField.Set(reflect.ValueOf(inst).Convert(instType))
	return nil
}
