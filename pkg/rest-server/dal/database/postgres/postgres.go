package postgres

import (
	"github.com/jmoiron/sqlx"
	// no-lint
	_ "github.com/lib/pq"

	db "github.com/kapustkin/go_guard/pkg/rest-server/dal/database"
)

// DB структура хранилища
type DB struct {
	db *sqlx.DB
}

func Init(conn string) *DB {
	connection, _ := sqlx.Connect("postgres", conn)
	return &DB{db: connection}
}

//GetParametrs
func (d *DB) GetParametrs() (*db.Parameters, error) {
	return &db.Parameters{}, nil
}

//GetAddressList
func (d *DB) GetAddressList() (*db.List, error) {
	return &db.List{}, nil
}
