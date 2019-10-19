package postgres

import (
	"fmt"
	"time"

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

type parametersTable struct {
	Create time.Time `db:"createdate"`
	K      int       `db:"k"`
	M      int       `db:"m"`
	N      int       `db:"n"`
}

//GetParametrs
func (d *DB) GetParametrs() (*db.Parameters, error) {
	parameters := []parametersTable{}
	err := d.db.Select(&parameters, `SELECT createDate,K,M,N FROM parameters ORDER by createDate DESC LIMIT 1`)
	if err != nil {
		return nil, err
	}
	if len(parameters) != 1 {
		return nil, fmt.Errorf("unexcepted rows count")
	}

	return &db.Parameters{
		Created: parameters[0].Create,
		K:       parameters[0].K,
		M:       parameters[0].M,
		N:       parameters[0].N}, nil
}

//GetAddressList
func (d *DB) GetAddressList() (*db.List, error) {
	return &db.List{}, nil
}
