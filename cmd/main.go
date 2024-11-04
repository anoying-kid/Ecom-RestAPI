package main

import (
	"database/sql"
	"log"

	"github.com/anoying-kid/Ecom/cmd/api"
	"github.com/anoying-kid/Ecom/config"
	"github.com/anoying-kid/Ecom/db"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = config.Envs.DBUser
	cfg.Passwd = config.Envs.DBPassword
	cfg.Addr = config.Envs.DBAddress
	cfg.DBName = config.Envs.DBName
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true

	db, err := db.NewMySQLStorage(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("DB: Successfully connected!")
}