package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

var myJwtSigningKey = []byte("jwt-secret")

/*
GenToken - A Util function to generate jwtToken which can be used in the request header
*/
func GenToken(id uint) (string, error) {

	/* Create Token */
	jwtToken := jwtgo.New(jwtgo.SigningMethodHS256)

	/* Set token claims */
	jwtToken.Claims = jwtgo.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * 3600).Unix(),
		"iss": "jwt-issuer",
	}

	/* Sign the token with our secret */
	return jwtToken.SignedString(myJwtSigningKey)
}
