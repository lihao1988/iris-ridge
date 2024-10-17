package service

import (
	"ridge/app/dao"
	"ridge/app/dto"
	"ridge/app/model"
	"ridge/common/lib/auth"
	"ridge/common/lib/db"

	"gorm.io/gorm"

	"github.com/kataras/iris/v12"
)

// UserService the service of user
type UserService struct {
	Service
}

// NewUserSvc new the instance of user_service
func NewUserSvc(db *gorm.DB, ctx iris.Context) *UserService {
	uSvc := &UserService{}
	uSvc.SetDB(db)
	uSvc.SetCtx(ctx)
	return uSvc
}

// UserLogin user login logic
func (uSvc *UserService) UserLogin() error {
	auth.SessionLogin(uSvc.GetCtx())
	return nil
}

// UserLogout user logout logic
func (uSvc *UserService) UserLogout() error {
	auth.SessionLogout(uSvc.GetCtx())
	return nil
}

// GetFirstUser get the first user_info
func (uSvc *UserService) GetFirstUser() (*dto.UserInfo, error) {
	cDao := dao.NewDao()
	uModel := model.NewUser(uSvc.GetDB())
	err := cDao.First(uModel)
	if err != nil {
		return nil, err
	}

	uInfo := &dto.UserInfo{}
	uInfo.Name = uModel.Name
	uInfo.Context = uModel.Context
	return uInfo, nil
}

// GetUserInfoById get user_info by user_id
func (uSvc *UserService) GetUserInfoById(id uint) (*dto.UserInfo, error) {
	uDao := dao.NewUserDao()
	uModel := model.NewUser(uSvc.GetDB())
	scopes := db.NewScopeOpts()
	scopes.Add(db.WithWhere("id = ?", id))
	err := uDao.Find(uModel, scopes.Export()...)
	if err != nil {
		return nil, err
	}

	uInfo := &dto.UserInfo{}
	uInfo.Name = uModel.Name
	uInfo.Context = uModel.Context
	return uInfo, nil
}
