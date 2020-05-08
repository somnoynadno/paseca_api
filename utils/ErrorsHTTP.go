package utils

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HandleBadRequest(w http.ResponseWriter, err error) {
	log.Warn(err)
	w.WriteHeader(http.StatusBadRequest)
	Respond(w, Message(false, err.Error()))
}

func HandleNotFound(w http.ResponseWriter) {
	log.Info("Not found")
	w.WriteHeader(http.StatusNotFound)
	Respond(w, Message(false, "Not found"))
}
