package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"time"
  "os"

	"github.com/golang-jwt/jwt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecryptPassword(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func EncryptPassword(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func GenerateJWT(secret, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func GenerateJwtEd25519(username string) (string, error) {

  privateKey, err := os.ReadFile("../../test_key.pem")
  if err != nil{
    return "", nil
  }
  
  
  key, err := jwt.ParseEdPrivateKeyFromPEM(privateKey)
  if err != nil {
    return "", fmt.Errorf(fmt.Sprintf("Problem formatting the ed25519 pem file %s", err.Error()))
  }


  token := jwt.New(jwt.SigningMethodEdDSA)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(10 * time.Minute)
  claims["authorized"] = true
  claims["user"] = username

  tokenString, err := token.SignedString(key)
  if err != nil {
    return "", err
  }

  return tokenString, nil
}
