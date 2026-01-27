package routes

import (
	"fintracker/internal/embeded"

	"github.com/pocketbase/pocketbase/apis"
)

var staticsHandler = apis.Static(embeded.Assets, true)
