package middleware

import (
	"errors"
	"net/http"
	u "paseca/utils"
)

var CheckAdminPermissions = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow admin path only to admins
		if r.Context().Value("context").(u.Values).Get("is_admin") == "false" {
			u.HandleForbidden(w, errors.New("restricted access"))
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
