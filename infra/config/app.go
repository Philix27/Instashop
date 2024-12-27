package config

import (
	"context"
	"instashop/db"

	"github.com/mikespook/gorbac"
)

type AppState struct {
	DbQueries *db.Queries
	Ctx       context.Context
	Rbac      *gorbac.RBAC
	Env       *AppEnv
}
