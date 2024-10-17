package abstract

import (
	"ridge/config/abstract/structs"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CGlobal global config
type CGlobal struct {
	Root      string // 根目录
	App       *structs.App
	Database  *structs.Database
	Migration *structs.Migration
}

// CGlobalAttrMap global "CGlobal" attribute
var CGlobalAttrMap map[string]string

// init config instance
func init() {
	// global "CGlobal" attribute
	var ok bool
	var tempFPName string
	CGlobalAttrMap = make(map[string]string)
	for fPName := range structs.RegConfigMap {
		tempFPName, ok = structs.AttrMap[fPName]
		if !ok {
			tempFPName = cases.Title(language.Und).String(fPName)
		}
		CGlobalAttrMap[fPName] = tempFPName
	}
}
