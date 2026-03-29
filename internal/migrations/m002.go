package migrations

import (
	_ "embed"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

//go:embed m002.up.json
var m002_up []byte

func m002() {
	m.Register(func(app core.App) error {
		return app.ImportCollectionsByMarshaledJSON(m002_up, false)
	}, func(app core.App) error {
		return nil
	})
}
