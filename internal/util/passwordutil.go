package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"login_app/internal/config"
	"login_app/internal/models"
	"time"
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
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["username"] = username

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJwtRSA(username string) (string, error) {

	block, res := pem.Decode(config.EnvConfigs.PRIVATE_SIGN_KEY)
	if block == nil {
		return "", fmt.Errorf(fmt.Sprintf("Problem decoding the ed25519 pem file %s", res))
	}

	if block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("invalid RSA private key")
	}

	file, err := os.ReadFile(config.EnvConfigs.PRIVATE_SIGN_KEY_PATH)
	if err != nil {
		return "", err
	}

	priv, err := jwt.ParseRSAPrivateKeyFromPEM(file)
	if err != nil {
		return "", err
	}

	claims := &models.JwtTokenCustomClaims{
		Username: username,
		Role:     "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(priv)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetPublicSignKey() ([]byte, error) {
	block, res := pem.Decode(config.EnvConfigs.PUBLIC_SIGN_KEY)
	if block == nil {
		return nil, fmt.Errorf(fmt.Sprintf("Problem decoding the ed25519 pem file %s", res))
	}
	return block.Bytes, nil
}
