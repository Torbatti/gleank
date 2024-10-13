package core

import (
	"database/sql"

	"github.com/Torbatti/gleank/models/settings"
	"github.com/nalgeon/redka"
)

type App interface {
	DataDir() string
	IsDev() bool

	Settings() *settings.Settings

	IsBootstrapped() bool
	Bootstrap() error

	DB() *sql.DB
	Store() *redka.DB
}
