package indoor

import (
	"ridge/common/global"
	"ridge/common/lib/auth"
)

// MiddleWare init middle ware
func MiddleWare() error {
	// session auth
	auth.InitSessions()

	// validate auth
	app := global.GApplication
	app.Use(auth.SessionValidate)

	return nil
}
