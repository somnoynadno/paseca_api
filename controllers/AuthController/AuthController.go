package AuthController

import (
	"encoding/json"
	"errors"
	"net/http"
	"paseca/models/auxiliary"
	u "paseca/utils"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &auxiliary.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	resp := auxiliary.Login(account.Email, account.Password)

	if resp["token"] == nil {
		u.HandleBadRequest(w, errors.New("wrong credentials"))
		return
	}

	u.Respond(w, resp)
}