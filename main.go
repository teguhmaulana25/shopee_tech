package main

import (
	"os"

	"github.com/teguhmaulana25/shopee_tech/config/database"
	"github.com/teguhmaulana25/shopee_tech/config/server"

	log "github.com/sirupsen/logrus"
)

func init() {
	//config env
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "teguhadmin123")
	os.Setenv("DBHOST", "localhost")
	os.Setenv("DBDATABASE", "db_shopee_exchange")
	os.Setenv("PORT", "8080")
	os.Setenv("ENV", "dev") // dev or production

	log.SetLevel(log.DebugLevel)
	if os.Getenv("ENV") == "dev" {
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	database.DbConnect()
	defer database.Db.Close()

	server.StartWebserver(os.Getenv("PORT"))
}
