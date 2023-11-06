package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	// TokenDuration is the duration of the token
	TokenDuration = 24 * time.Hour
	Secret        = "so_secret!"
)

/**
 * This function generates a token.
 * @param mail The email of the token.
 * @return the token if it was generated, an error otherwise.
 */
func GenerateTokenFromMail(mail string) (string, error) {
	key := []byte(Secret)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"sub": mail,
		"exp": time.Now().Add(TokenDuration).Unix(), // Set the expiration time
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/**
 * This function verifies a token.
 * @param tokenString The token to verify.
 * @return the email of the token if it was verified, an error otherwise.
 */
func GetMailFromToken(tokenString string) (string, error) {
	key := []byte(Secret)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		return "", err // If the token expires, it will return an error here
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		subject := claims["sub"].(string)
		return subject, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}

/**
 * This function hashes a password.
 * @param plain_password The password to hash.
 * @return the hashed password.
 */
func HashPassword(plain_password string) string {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(plain_password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed_password)
}

/**
 * This function verifies a password.
 * @param hashed_password The hashed password.
 * @param plain_password The plain password.
 * @return true if the password is correct, false otherwise.
 */
func VerifyPassword(hashed_password string, plain_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(plain_password))
	return err != nil
}
