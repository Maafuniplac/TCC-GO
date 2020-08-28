package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver Postgres para Go
)

//Funçao para conectar no bando de dados
func ConectaComDB() *sql.DB {
	conexao := "user=postgres dbname=Notifica_Go password=Maafdrums1 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
