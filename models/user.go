package models

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Todos     *[]Todo   `json:"todos"`
}

// HashPassword encrypt password
func (u *User) HashPassword() error {
	passBytes := []byte(u.Password)
	hashBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashBytes)
	return nil
}

// CheckPassword verifies if password is correct
func (u *User) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}

// GenerateToken generates jwt token
func (u *User) GenerateToken() (*Auth, error) {
	expiresAt := time.Now().Add(12 * time.Hour).Unix()
	claims := jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    "basicApp",
		Id:        strconv.Itoa(u.ID),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("JWT_KEY")
	token, err := accessToken.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &Auth{
		Token: token,
		User:  u,
	}, nil
}
