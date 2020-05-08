package utils

import (
	"encoding/json"
	"net/http"
)

func SetTotalCountHeader(w http.ResponseWriter, count string) {
	w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
	w.Header().Add("X-Total-Count", count)
}

func CheckOrderAndSortParams(order *string, sort *string) {
	if *order != "ASC" && *order != "DESC" {
		*order = "ASC"
	}
	if *sort == "" {
		*sort = "ID"
	}
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func RespondJSON(w http.ResponseWriter, data []byte) {
	_, _ = w.Write(data)
}

var HandleOptions = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
