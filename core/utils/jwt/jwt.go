package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func InitJWT() (*Authorization, error) {
	root := Authorization{}
	err := root.initJWT()
	return &root, err
}

func (root *Authorization) initJWT() error {
	root.JwtSecret = []byte("rK7fL9mN3pQ5sT8vU0wX2yB4zC6dE7gF1hJ3kM5nO8pR0sT2uV4xW6yZ8")
	root.JwtExpiry = time.Hour * 24
	return nil
}

func (root *Authorization) GenJWT(UserID string) (string, error) {
	expireTime := time.Now().Add(root.JwtExpiry)

	claims := &jwt.MapClaims{
		"sub": UserID,
		"exp": expireTime.Unix(),
		"iat": time.Now().Unix(),
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := genToken.SignedString(root.JwtSecret)

	return token, err
}

func (root *Authorization) RecoverData(rawToken string) (string, bool) {
	token, err := jwt.Parse(rawToken, root.checkSigningMethod)

	if err != nil || !token.Valid {
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		userinfo, ok := (claims)["sub"].(string)
		if ok {
			return userinfo, true
		}
	}

	return "", false
}

func (root *Authorization) checkSigningMethod(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		fmt.Printf("unsupport SigningMethod: %v\n", token.Header["alg"])
		return nil, fmt.Errorf("unsupport SigningMethod: %v", token.Header["alg"])
	}

	return root.JwtSecret, nil
}
