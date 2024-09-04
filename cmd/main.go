package main

import (
	"github.com/brianleogoldman/goshop/cmd/api"
	"github.com/brianleogoldman/goshop/config"
	database "github.com/brianleogoldman/goshop/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := database.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
