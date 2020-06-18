package middleware

import (
	"errors"
	"net/http"
	u "paseca/utils"
	"strconv"
	"strings"
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

var CheckPermissionsBySubscription = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path // path of current request
		subscriptionType, _ := strconv.ParseInt(r.Context().Value("context").(u.Values).Get("subscription_type_id"), 0, 64)

		var isExpired bool
		if r.Context().Value("context").(u.Values).Get("subscription_expired") == "false" {
			isExpired = false
		} else {
			isExpired = true
		}

		permitPath := func (prefix string, permittedID int64) int {
			if strings.HasPrefix(path, prefix) {
				if subscriptionType < permittedID {
					return 0
				} else if isExpired {
					return 0
				}
			}
			return 1
		}

		allowed := 1
		allowed &= permitPath("/api/lk/honey_harvest", 2)
		allowed &= permitPath("/api/lk/honey_sale", 2)
		allowed &= permitPath("/api/lk/control_harvest", 2)
		allowed &= permitPath("/api/lk/wiki", 2)
		allowed &= permitPath("/api/lk/pollen_harvest", 3)
		allowed &= permitPath("/api/lk/family_disease", 3)
		allowed &= permitPath("/api/lk/swarm", 3)
		allowed &= permitPath("/api/lk/honey_harvest_stats", 4)

		if allowed != 1 {
			if isExpired {
				u.HandlePaymentRequired(w, errors.New("subscription expired"))
			} else {
				u.HandlePaymentRequired(w, errors.New("subscription not sufficient"))
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
