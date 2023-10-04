package infraestructure

import (
	"database/sql"
	"fmt"
	config "golang-template-clean-architecture/src/common/config"
	"log"

	_ "github.com/lib/pq"
)
type DBConnection struct{
	*sql.DB
}

func NewDBConnection(cfg * config.Config) *DBConnection{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Dbname,
	)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error to connect database : %s", err)
	}
	return &DBConnection{result}
}