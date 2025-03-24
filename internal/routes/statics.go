package routes

import (
	"fintracker/internal/embeded"

	"github.com/pocketbase/pocketbase/apis"
)

// var staticsHandler = filesystem.New(filesystem.Config{
// 	Root:       ,
// 	PathPrefix: "statics",
// 	MaxAge:     60 * 60,
// })

var assetsHandler = apis.Static(embeded.Assets, true)
