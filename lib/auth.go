package lib

import (
	"errors"
	"time"

	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/dgrijalva/jwt-go"
	"github.com/kejarmimpi/models"
)

/*
 Get authenticated user and update token
*/
func Authenticate(email string, password string) (user *models.Users, err error) {
	msg := "invalid email or password."
	user = &models.Users{Email: email}
	SecretKey := []byte("something")

	if err := user.Read("Email"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		// No user
		return user, errors.New(msg)
	} else if user.Password != convert.StrTo(password).Md5() {
		// No matched password
		return user, errors.New(msg)
	} else {
		// Create JWT token
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			token.Claims = jwt.MapClaims{
		        "exp": time.Now().Add(time.Hour * 6).Unix(),
		        "id": user.Id,
		    }
		    tokenString ,_ := token.SignedString(SecretKey)

		    user.Token = tokenString
			
		user.Update("Token")
		return user, nil
	}
}