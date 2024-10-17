package dto

import "ridge/common/dto"

// LoginReq login request
type LoginReq struct {
	Name     string `json:"name" validate:"required,min=6" label:"用户名"`
	Password string `json:"password" validate:"required,min=8" label:"密码"`
}

// UserInfoReq userInfo request
type UserInfoReq struct {
	Name string `json:"name" url:"name" validate:"required" label:"用户名"`
}

// CreateUserReq create user request
type CreateUserReq struct {
	Name    string `json:"name" form:"name" validate:"required,min=6" label:"用户名称"`
	Context string `json:"context" form:"context" validate:"required,min=8" label:"备注信息"`
}

type UserInfoResp struct {
	dto.Response
	Data UserInfo `json:"data"`
}

// UserInfo user information
type UserInfo struct {
	Name    string `json:"name" label:"用户名称"`
	Context string `json:"context" label:"备注信息"`
}
