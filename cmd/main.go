package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/yann-fk-21/todo-platform/cmd/api"
	"github.com/yann-fk-21/todo-platform/config"
	"github.com/yann-fk-21/todo-platform/db"
)

const ADDRESS = ":8000"

func main() {
    
	db, err := db.NewMySQLStorage(
		mysql.Config{
			User: config.Envs.DBUser,
			Passwd: config.Envs.DBPassword,
			Net: "tcp",
			Addr: config.Envs.DBAddress,
			DBName: config.Envs.DBName,
			AllowNativePasswords: true,
			ParseTime: true,
		})

		if err != nil {
			log.Fatal(err)
		}

		initStorage(db)

	s := api.NewApiServer(ADDRESS, db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connect successfully")
}