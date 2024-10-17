package dao

import (
	baseDao "ridge/common/dao"
)

// Dao the dao's instance of app
type Dao struct {
	baseDao.Dao
}

// NewDao new the instance of "dao"
func NewDao() *Dao {
	dao := &Dao{}
	return dao
}
