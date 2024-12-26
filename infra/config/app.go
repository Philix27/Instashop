package config

import (
	"context"
	"instashop/db"
)

type AppState struct {
	DbQueries *db.Queries
	Ctx       context.Context
}
