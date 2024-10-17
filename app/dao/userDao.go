package dao

import (
	"ridge/common/lib/db"
	"ridge/common/model"

	"gorm.io/gorm"
)

// UserDao the dao of user
type UserDao struct {
	Dao
}

// NewUserDao new the instance of user_dao
func NewUserDao() *UserDao {
	return &UserDao{}
}

// FindOne get one record for 'option'
func (ud *UserDao) FindOne(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return db.FindOne(m, options...)
}
