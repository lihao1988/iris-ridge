package autoload

import (
	"fmt"
	"testing"

	"ridge/common/global"
	"ridge/common/share/maintesting"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestAutoload(t *testing.T) {
	app := App{}
	err := app.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(global.GConfig)
	fmt.Println(global.GConfig.App)
	fmt.Println(global.GConfig.Database.Mysql)
}

func TestZero(t *testing.T) {
	str := "mysql"
	newTitle := cases.Title(language.Und).String(str)
	fmt.Println(newTitle)
}

func TestMain(m *testing.M) {
	maintesting.Run(m, "/home/anyuan/lgm_study/ridge")
}
