package auth

import (
	"fmt"
	"time"

	"ridge/common/global"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const cookieNameForSessionID = "session_id"

// InitSessions init session
func InitSessions() {
	global.SessionManager = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires: 30 * time.Minute,
		// CookieSecureTLS: true,
		AllowReclaim:                true,
		DisableSubdomainPersistence: true,
	})

	app := global.GApplication
	app.Use(global.SessionManager.Handler())
}

// SessionValidate session auth validate
func SessionValidate(ctx iris.Context) {
	session := sessions.Get(ctx)
	fmt.Println(ctx.GetCookie(cookieNameForSessionID))
	fmt.Println(session.GetBoolean("authenticated"))

	ctx.Next()
}

// SessionLogin set session for login
func SessionLogin(ctx iris.Context) {
	// 创建或获取当前会话
	session := global.SessionManager.Start(ctx)

	// Set user as authenticated
	session.Set("authenticated", true)
}

// SessionLogout login out
func SessionLogout(ctx iris.Context) {
	session := sessions.Get(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
}
