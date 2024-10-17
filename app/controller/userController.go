package controller

import (
	"ridge/app/dto"
	"ridge/app/service"
	"ridge/common/global"
	"ridge/common/lib/request"
)

// UserController controller user
type UserController struct {
	Controller
}

// init append "user" to GControllers
func init() {
	// add register controller
	global.AddControllerRegister(new(UserController))
}

// Prefix custom routing prefix of "user" controller
func (u *UserController) Prefix() string {
	return "user"
}

// BuildRoute build the controller route
func (u *UserController) BuildRoute() {
	u.Mvc.Handle(u)
}

// PostLogin user login operate
// @Tags 用户功能模块
// @Summary 用户登录
// @Description 用于用户登录操作
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "auth token"
// @Param data body dto.LoginReq true "用户信息"
// @Success 200 {object} dto.APISuccess	"请求成功"
// @Failure 400 {object} dto.APIError "请求错误"
// @Router /app/user/login [post]
func (u *UserController) PostLogin() {
	// request data
	loginReq := &dto.LoginReq{}
	if !request.ValidReadJson(u.Ctx, loginReq) {
		return
	}

	// service
	usVc := service.NewUserSvc(u.GetDB(), u.Ctx)
	err := usVc.UserLogin()
	u.Send(nil, err)
	return
}

// GetLogout user logout operate
// @Tags 用户功能模块
// @Summary 用户登出
// @Description 用于用户登出操作
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "auth token"
// @Success 200 {object} dto.APISuccess	"请求成功"
// @Failure 400 {object} dto.APIError "请求错误"
// @Router /app/user/logout [get]
func (u *UserController) GetLogout() {
	usVc := service.NewUserSvc(u.GetDB(), u.Ctx)
	err := usVc.UserLogout()
	u.Send(nil, err)
	return
}

// PostInfo create user info
// @Tags 用户功能模块
// @Summary 创建新用户
// @Description 用于创建新用户
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "auth token"
// @Param name formData string true "用户名称"
// @Param context formData string true "备注信息"
// @Success 200 {object} dto.APISuccess	"请求成功"
// @Failure 400 {object} dto.APIError "请求错误"
// @Router /app/user/info [post]
func (u *UserController) PostInfo() {
	uInfoReq := &dto.CreateUserReq{}
	if !request.ValidReadForm(u.Ctx, uInfoReq) {
		return
	}

	u.Send(uInfoReq, nil)
	return
}

// GetInfo demo function
// @Tags 用户功能模块
// @Summary 获取用户信息
// @Description 用于获取用户信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "auth token"
// @Param data query dto.UserInfoReq true "用户名称"
// @Success 200 {object} dto.APISuccess	"请求成功"
// @Failure 400 {object} dto.APIError "请求错误"
// @Router /app/user/info [get]
func (u *UserController) GetInfo() {
	uInfoReq := &dto.UserInfoReq{}
	if !request.ValidReadQuery(u.Ctx, uInfoReq) {
		return
	}

	usVc := service.NewUserSvc(u.GetDB(), u.Ctx)
	uInfo, err := usVc.GetFirstUser()
	u.Send(uInfo, err)
	return
}
