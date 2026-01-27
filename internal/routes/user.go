package routes

import (
	"fintracker/internal/controllers"
	"fintracker/internal/urls"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

var userIdStoreKey = "user_id"

func basicAuthMiddleware(c controllers.Controllers, e *core.RequestEvent) error {
	basicUnathorized := func() error {
		e.Response.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		return e.UnauthorizedError("Unauthorized", "")
	}

	email, pwd, ok := e.Request.BasicAuth()
	if !ok {
		return basicUnathorized()
	}

	userId, err := c.User.GetUserIdFromCredentials(email, pwd)
	if err != nil {
		return basicUnathorized()
	}

	e.Set(userIdStoreKey, userId)

	return e.Next()
}

func assetScopeCheckMiddleware(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)
	userId := e.Get(userIdStoreKey).(string)

	fmt.Println(assetId, userId)
	if err := c.User.ChechUserOwnsAsset(userId, assetId); err != nil {
		return e.NotFoundError("asset not found", "")
	}

	return e.Next()
}

func priceScopeCheckMiddleware(c controllers.Controllers, e *core.RequestEvent) error {
	priceId := e.Request.PathValue(urls.PriceIdPathParam)
	userId := e.Get(userIdStoreKey).(string)

	if err := c.User.ChechUserOwnsPrice(userId, priceId); err != nil {
		return e.NotFoundError("price not found", "")
	}

	return e.Next()
}
