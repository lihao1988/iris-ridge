package model

import (
	"ridge/common/global"

	"gorm.io/gorm"
)

// UserModel model demo
type UserModel struct {
	Model
	Name    string
	Context string
}

// init append "User" to GModels
func init() {
	// add register model
	global.AddModelRegister(new(UserModel))
}

// NewUser new user model
func NewUser(db *gorm.DB) *UserModel {
	m := &UserModel{}
	m.SetDB(db)
	return m
}

// TableName get the table name
func (u *UserModel) TableName() string {
	return "user"
}

// TableComment table comment
func (u *UserModel) TableComment() string {
	return "用户信息表"
}
