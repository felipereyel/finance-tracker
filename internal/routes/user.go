package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func basicAuthMiddleware(e *core.RequestEvent) error {
	basicUnathorized := func() error {
		e.Response.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		return e.Error(http.StatusUnauthorized, "Unauthorized", "")
	}

	username, pwd, ok := e.Request.BasicAuth()
	if !ok {
		fmt.Println("basic auth not ok")
		return basicUnathorized()
	}

	user, err := e.App.FindAuthRecordByEmail("users", username)
	if err != nil {
		fmt.Println("auth record not found", err)
		return basicUnathorized()
	}

	if !user.ValidatePassword(pwd) {
		fmt.Println("pwd not matched")
		return basicUnathorized()
	}

	e.Set("user_id", user.Id)

	return e.Next()
}
