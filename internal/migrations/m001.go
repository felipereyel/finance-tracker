package migrations

import (
	_ "embed"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

//go:embed m001.up.json
var m001_up []byte

func m001() {
	m.Register(func(app core.App) error {
		return app.ImportCollectionsByMarshaledJSON(m001_up, false)
	}, func(app core.App) error {
		return nil
	})
}
