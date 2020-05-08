package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"paseca/db"
	"paseca/middleware"
	u "paseca/utils"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	router := mux.NewRouter()

	// Handle options
	router.HandleFunc("/api/{any}", u.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/api/{any1}/{any2}", u.HandleOptions).Methods("OPTIONS")

	// middleware usage
	// do NOT modify the order
	router.Use(middleware.LogBody) // log HTTP body, method and URI
	router.Use(middleware.CORS)   // enable CORS headers

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

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Panic(err)
	}
}
