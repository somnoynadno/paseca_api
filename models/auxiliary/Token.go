package auxiliary

import "github.com/dgrijalva/jwt-go"

/*
JWT claims struct
*/

type Token struct {
	UserID  uint
	IsAdmin bool
	SubscriptionTypeID  uint
	SubscriptionExpired bool
	jwt.StandardClaims
}
