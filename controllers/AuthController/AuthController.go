package AuthController

import (
	"encoding/json"
	"errors"
	"net/http"
	"paseca/db"
	"paseca/models"
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

type NewAccountRequest struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
}

var Registration = func(w http.ResponseWriter, r *http.Request) {
	account := &NewAccountRequest{}
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Where("email = ?", account.Email).First(&user).Error

	if err == nil {
		u.HandleBadRequest(w, errors.New("user with this email already exist"))
		return
	}

	user.Email = account.Email
	user.Password = account.Password
	user.Surname = account.Surname
	user.Name = account.Name

	user.SubscriptionTypeID = 1
	user.SubscriptionStatusID = 1
	user.IsAdmin = false

	err = db.Create(&user).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}