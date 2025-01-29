package services

import (
	"context"
	"fmt"
	"ladipage_server/apis/entities"
	"ladipage_server/common/configs"
	customerrors "ladipage_server/core/custom_errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	config *configs.Configs
}

func NewJwtService(config *configs.Configs) *JwtService {
	return &JwtService{
		config: config,
	}
}

func (u *JwtService) GenToken(ctx context.Context, userName string, userId int64, updatedAt time.Time) (*entities.JwtResponse, error) {
	expirationDuration, err := time.ParseDuration(u.config.ExpireAccess)
	if err != nil {
		return nil, fmt.Errorf("invalid expiration duration: %v", err)
	}

	claims := entities.User{
		UserName:           userName,
		Id:                 userId,
		UpdatedAccountUser: updatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.config.AccessSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %v", err)
	}

	return &entities.JwtResponse{
		Token:              tokenString,
		UserName:           userName,
		UserId:             userId,
		UpdatedAccountUser: updatedAt,
	}, nil
}

func (u *JwtService) VerifyToken(ctx context.Context, tokenString string) (*entities.User, *customerrors.CustomError) {
	token, err := jwt.ParseWithClaims(tokenString, &entities.User{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.config.AccessSecret), nil
	})

	if err != nil {
		return nil, customerrors.ErrVerifyToken
	}

	if claims, ok := token.Claims.(*entities.User); ok && token.Valid {
		return claims, nil
	}

	return nil, customerrors.ErrVerifyToken
}
