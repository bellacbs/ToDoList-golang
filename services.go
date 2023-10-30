package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var JWTsecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GenerateToken(user *UserResponseDTO) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	fmt.Println(time.Now().Add(24 * time.Hour))
	tokenString, err := token.SignedString(JWTsecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (*UserResponseDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTsecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	userMap := claims["user"]
	data, ok := userMap.(map[string]interface{})
	if !ok {
		fmt.Println("Error to convert map")
	}
	uuidValue, err := uuid.Parse(data["id"].(string))
	if err != nil {
		return nil, fmt.Errorf("Error To convert String to uuid")
	}
	user := UserResponseDTO{ID: uuidValue, Email: data["email"].(string), Name: data["name"].(string)}
	expirationTime := int64(claims["exp"].(float64))
	currentTime := time.Now().Unix()
	if currentTime > expirationTime {
		return nil, fmt.Errorf("Expired Token")
	}
	return &user, nil
}
