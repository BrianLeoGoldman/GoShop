package main

import (
	"github.com/brianleogoldman/goshop/cmd/api"
	database "github.com/brianleogoldman/goshop/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := database.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "pass123",
		Addr:                 "127.0.0.1:3306",
		DBName:               "goshop",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
