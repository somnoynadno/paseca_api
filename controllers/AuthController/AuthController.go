package AuthController

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
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
	Captcha   string `json:"captcha"`
}

type CaptchaResponse struct {
	Success bool `json:"success"`
}

var Registration = func(w http.ResponseWriter, r *http.Request) {
	account := &NewAccountRequest{}
	user := models.User{}
	cr := &CaptchaResponse{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// verify captcha
	secret := os.Getenv("captcha_secret")
	if secret == "" {
		panic("no captcha secret provided")
	}

	resp, err := http.Get("https://www.google.com/recaptcha/api/siteverify?" +
		"secret=" + secret + "&" +
		"response=" + account.Captcha)

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, cr)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	if cr.Success == false {
		u.HandleUnauthorized(w, errors.New("bad captcha"))
		return
	}

	db := db.GetDB()
	err = db.Where("email = ?", account.Email).First(&user).Error

	if err == nil {
		u.HandleBadRequest(w, errors.New("user with this email already exist"))
		return
	}

	hash, err := u.HashPassword(account.Password)
	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	user.Password = hash
	user.Email = account.Email
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