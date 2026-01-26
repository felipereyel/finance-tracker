package routes

import (
	"fintracker/internal/components"
	"fintracker/internal/controllers"
	"fintracker/internal/models"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func setupScopedPricesRoutes(group *router.RouterGroup[*core.RequestEvent], c controllers.Controllers) {
	// TODO: add scope check for price_id
	// group.BindFunc(basicAuthMiddleware)

	group.GET("/", withControllerClousure(c, priceDetails))
	group.POST("/", withControllerClousure(c, priceUpdate))
}

func priceDetails(c controllers.Controllers, e *core.RequestEvent) error {
	priceId := e.Request.PathValue("price_id")

	price, err := c.Price.GetPrice(priceId)
	if err != nil {
		fmt.Println("Error retrieving price:", err)
		return err
	}

	return sendPage(e, components.PricePage(price))
}

func priceUpdate(c controllers.Controllers, e *core.RequestEvent) error {
	priceId := e.Request.PathValue("price_id")

	priceDTO := models.PriceUpdateDTO{}
	if err := e.BindBody(&priceDTO); err != nil {
		return err
	}

	if err := c.Price.UpdatePrice(priceId, priceDTO); err != nil {
		return e.JSON(500, map[string]any{"error": true})
	}

	return e.JSON(200, map[string]any{"success": true})
}
