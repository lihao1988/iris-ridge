package db

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestInitOracle(t *testing.T) {
	gormDb, err := InitGormOracle()
	if err != nil {
		panic(err)
	}

	type Test struct {
		gorm.Model
		id uint
	}

	var test Test
	err = gormDb.Table("TESTS").Unscoped().Find(&test, "ID = ?", 2).Error
	fmt.Println(test, err)

	type Users struct {
		gorm.Model
		id uint
		tx string
	}
	uModel := &Users{}
	uList := []Users{}
	err = gormDb.Model(uModel).Unscoped().Find(&uList).Error
	fmt.Println(uList, err)
}
