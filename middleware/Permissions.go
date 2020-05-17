package middleware

import (
	"errors"
	"net/http"
	u "paseca/utils"
	"strings"
)

var CheckPermissions = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do NOT check OPTIONS
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		// allow everybody to auth
		if strings.HasPrefix(r.URL.Path, "/api/auth") {
			next.ServeHTTP(w, r)
			return
		}

		// allow admin path only to admins
		if strings.HasPrefix(r.URL.Path, "/api/admin") &&
			r.Context().Value("context").(u.Values).Get("is_admin") == "false" {
				u.HandleForbidden(w, errors.New("restricted access"))
				return
		}

		next.ServeHTTP(w, r)
	})
}
