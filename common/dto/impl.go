package dto

// Response common response
type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data" label:"用户信息"`
	Msg  string      `json:"msg" example:"success"`
	// TraceId string      `json:"traceId" example:"12123213213"`
}
