package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	jwt.RegisteredClaims
	Id                 int64     `json:"id,omitempty"`
	UserName           string    `json:"user_name,omitempty"`
	UpdatedAccountUser time.Time `json:"updated_at,omitempty"`
}

type JwtResponse struct {
	Token              string    `json:"token"`
	UserName           string    `json:"user_name"`
	UserId             int64     `json:"user_id"`
	UpdatedAccountUser time.Time `json:"updated_at,omitempty"`
}

type LoginResponse struct {
	Token     string    `json:"token,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type GoogleClaims struct {
	Iss           string `json:"iss"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"` //gg ID
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	jwt.RegisteredClaims
}
