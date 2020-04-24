package hhttp

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type authToken struct {
	Token string `json:"token,omiempty"`
	Expired int64 `json:"token,omiempty"`
}

type authClaims struct {
	UserID int `json:userId`
	jwt.StandardClaims
}

func verify(token string) error {
	_, err := parseJWTToken(token)
	if err != nil {
		return err
	}

	return nil
}

func isExprired(err error) bool {
	if e, ok := err.(*jwt.ValidationError); ok && e.Errors == jwt.ValidationErrorExpired {
		return true
	}

	return false
}


func parseJWTToken(token string) (*jwt.Token, error){
	key := viper.GetString("key")
	return jwt.ParseWithClaims(token, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}