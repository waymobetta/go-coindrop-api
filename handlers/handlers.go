package handlers

import "github.com/waymobetta/go-coindrop-api/db"

// Handlers ...
type Handlers struct {
	db *db.DB
}

// Config ...
type Config struct {
	DB *db.DB
}

// NewHandlers ...
func NewHandlers(config *Config) *Handlers {
	return &Handlers{
		db: config.DB,
	}
}
