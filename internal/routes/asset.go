package routes

import (
	"fintracker/internal/components"
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"

	"github.com/pocketbase/pocketbase/core"
)

func homeRedirect(e *core.RequestEvent) error {
	return e.Redirect(302, "/assets")
}

func assetList(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	wallet := e.Request.URL.Query().Get("wallet")

	aggregated, err := db.GetCurrentSummary(wallet)
	if err != nil {
		return err
	}

	summary := models.Summary{
		Aggregates: make([]models.AssetAggregate, 0),
	}

	for _, asset := range aggregated {
		summary.Total += asset.LastPrice
		summary.Aggregates = append(summary.Aggregates, asset)
	}

	return sendPage(e, components.AssetSummaryPage(summary))
}

// func taskNew(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	task, err := tc.CreateTask(user.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.Redirect("/edit/" + task.Id)
// }

// func taskEdit(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	taskId := c.Params("id")
// 	task, err := tc.RetrieveTask(user.ID, taskId)

// 	if err != nil {
// 		return err
// 	}

// 	return sendPage(c, components.TaskEditPage(task))
// }

// func taskSave(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	var taskId = c.Params("id")
// 	var taskChange controllers.TaskChange
// 	err := c.BodyParser(&taskChange)
// 	if err != nil {
// 		return err
// 	}

// 	if err := tc.UpdateTask(user.ID, taskId, taskChange); err != nil {
// 		return err
// 	}

// 	return c.SendStatus(fiber.StatusOK)
// }
