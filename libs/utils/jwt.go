package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go-clean-architure/libs/common"
	"io/ioutil"
	"os"
	"time"
)

func GenerateToken(data common.JSON) (string, error) {

	//  token is valid for 7days
	date := time.Now().Add(time.Hour * 24)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}
