package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersDbUsername = "user_db_username"
	usersDbPassword = "user_db_password"
	usersDbHost     = "user_db_host"
	usersDbSchema   = "user_db_schema"
)

var (
	//Database conexão com o banco de dados
	Database *sql.DB
	erro     error
	username = os.Getenv(usersDbUsername)
	password = os.Getenv(usersDbPassword)
	host     = os.Getenv(usersDbHost)
	schema   = os.Getenv(usersDbSchema)
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)

	Database, erro = sql.Open("mysql", dataSourceName)

	if erro != nil {
		panic(erro)
	}

	if erro = Database.Ping(); erro != nil {
		panic(erro)
	}

	log.Println("Conexão com banco de dados realizada com sucesso")

}
