package routes

import (
	"fintracker/internal/components"

	"github.com/pocketbase/pocketbase/core"
)

func notFoundHandler(e *core.RequestEvent) error {
	return sendPage(e, components.NotFoundPage())
}
