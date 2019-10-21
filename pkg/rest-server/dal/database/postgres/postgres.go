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

type addressListsTable struct {
	Create  time.Time `db:"createdate"`
	IsWhite bool      `db:"iswhite"`
	Network string    `db:"network"`
}

//GetParametrs
func (d *DB) GetParametrs() (*db.Parameters, error) {
	if d.db == nil {
		return nil, fmt.Errorf("no connection to database")
	}
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
func (d *DB) GetAddressList() (*[]db.List, error) {
	if d.db == nil {
		return nil, fmt.Errorf("no connection to database")
	}

	rawData := []addressListsTable{}
	err := d.db.Select(&rawData, `SELECT createDate,IsWhite, Network FROM addressLists ORDER by Id DESC`)
	if err != nil {
		return nil, err
	}

	var result = make([]db.List, len(rawData))
	for i, r := range rawData {
		result[i] = db.List{
			Created: r.Create,
			IsWhite: r.IsWhite,
			Network: r.Network,
		}
	}

	return &result, nil
}
