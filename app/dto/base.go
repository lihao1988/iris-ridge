package dto

import "ridge/common/dto"

// APISuccess api request success
type APISuccess struct {
	dto.Response
}

// APIError api request error
type APIError struct {
	dto.Response
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"error"`
}
