package auxiliary

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"time"
)

type Account struct {
	Email    string
	Password string
}

func Login(email, password string) map[string]interface{} {
	user := &models.User{}
	err := db.GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		log.Warn(err)
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if user.Password != password { // Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	// Worked! Logged In
	db.GetDB().Model(&user).Update("LastLogin", time.Now())

	// check if subscription is expired
	IsExpired := false
	if user.SubscriptionEnd != nil && user.SubscriptionEnd.Before(time.Now()) {
		IsExpired = true
		db.GetDB().Model(&user).Update("SubscriptionStatusID", 2)
	} else {
		db.GetDB().Model(&user).Update("SubscriptionStatusID", 1)
	}

	// Create JWT token
	tk := &Token{UserID: user.ID, IsAdmin: user.IsAdmin,
		SubscriptionTypeID: user.SubscriptionTypeID, SubscriptionExpired: IsExpired}

	tk.ExpiresAt = time.Now().Add(time.Hour * 72).Unix() // valid for 3 days
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Logged In")
	resp["token"] = tokenString
	resp["user_id"] = user.ID
	resp["subscription_type_id"] = user.SubscriptionTypeID
	resp["subscription_expired"] = IsExpired

	return resp
}