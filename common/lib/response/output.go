package response

import "github.com/kataras/iris/v12"

// Msg output in json
func Msg(ctx iris.Context, code int, data interface{}, msg string) {
	// msg - default value
	if msg == "" {
		switch code {
		case iris.StatusOK:
			msg = "success"
		default:
			msg = "error"
		}
	}

	// response data
	res := map[string]interface{}{
		"code": code,
		"data": data,
		"msg":  msg,
		// "traceId": "", // not yet used
	}
	_ = ctx.JSON(res)
}

// SendMsg output of data
func SendMsg(ctx iris.Context, data interface{}, err error, msg ...string) {
	if err != nil {
		Msg(ctx, iris.StatusBadRequest, nil, err.Error())
		return
	}

	var msgStr string
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	Msg(ctx, iris.StatusOK, data, msgStr)
	return
}
