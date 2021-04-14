package genkey

import (
	"crypto/rsa"

	"io/ioutil"
	"log"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

func GenPublicKey(pubPath string) *rsa.PublicKey {
	key, err := ioutil.ReadFile(pubPath)
	if err != nil {
		log.Fatal(err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		log.Fatal(err)
	}
	return pubKey
}
