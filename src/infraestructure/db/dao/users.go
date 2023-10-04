package infraestructure

import (
	"database/sql"
	db "golang-template-clean-architecture/src/infraestructure/db/adapter"
)

type PostgreUserDao struct {
	sql *sql.DB
}

func NewPostgreUserDao(connection *db.DBConnection) *PostgreUserDao{
	return &PostgreUserDao{ sql: connection.DB}
}