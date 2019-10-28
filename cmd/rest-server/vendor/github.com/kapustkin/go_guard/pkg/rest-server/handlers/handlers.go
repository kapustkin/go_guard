package handlers

import (
	"github.com/kapustkin/go_guard/pkg/rest-server/dal/database"
	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
)

type MainHandler struct {
	store storage.Storage
	db    database.Database
}

// Init main handler
func Init(st *storage.Storage, db *database.Database) *MainHandler {
	return &MainHandler{store: *st, db: *db}
}
