package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"paseca/db"
	"paseca/router"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	Router := router.InitRouter()

	// check connection
	con := db.GetDB()
	errors := con.GetErrors()
	if errors != nil && len(errors) > 0 {
		panic(errors[0])
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060" // localhost
	}

	log.Info("Listening on: ", port)

	err := http.ListenAndServe(":" + port, Router)
	if err != nil {
		log.Panic(err)
	}
}
