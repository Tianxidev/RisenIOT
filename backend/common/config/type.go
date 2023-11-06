package config

import (
	"net/http"
)

type Conf interface {
	SetDb(db *DBConfig)
	GetDb() *DBConfig

	SetEngine(engine http.Handler)
	GetEngine() http.Handler
}
