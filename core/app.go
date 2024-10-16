package core

import (
	"context"
	"database/sql"

	"github.com/Torbatti/gleank/models/settings"
	"github.com/go-chi/jwtauth/v5"
	"github.com/nalgeon/redka"
)

type App interface {
	DataDir() string
	IsDev() bool

	Settings() *settings.Settings

	IsBootstrapped() bool
	Bootstrap() error

	DB() *sql.DB
	DB_Context() context.Context
	Store() *redka.DB

	TokenAuth() *jwtauth.JWTAuth
}
