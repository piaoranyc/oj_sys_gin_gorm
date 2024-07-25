package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("gingormkey")

func TestGenJwt(t *testing.T) {
	userClaim := UserClaims{
		Identity:       "123456",
		Name:           "get",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Error(err)

	}
	fmt.Println(tokenString)
}

func TestAnalyseToken(t *testing.T) {

}
