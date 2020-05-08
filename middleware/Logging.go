package middleware

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var LogBody = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("Request from %s on %s with %s method", r.Host, r.RequestURI, r.Method))

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Error reading body: %v", err)
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}

		if len(body) > 0 {
			log.Debug(string(body))
		}

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}
