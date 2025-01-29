package domain

import (
	"context"
	"gorm.io/gorm"
	"ladipage_server/apis/entities"
)

type Users struct {
	*entities.Model
	UserName     string `json:"user_name,omitempty" gorm:"column:user_name;unique"`
	Password     string `json:"password,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	GoogleUserId string `json:"google_user_id,omitempty" gorm:"column:google_user_id;unique"`
	Email        string `json:"email,omitempty"`
	NickName     string `json:"nick_name,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Role         int32  `json:"role,omitempty"`
}

type RepositoryUser interface {
	Create(ctx context.Context, db *gorm.DB, user *Users) error
	Update(ctx context.Context, user *Users) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*Users, error)
	FindByUsername(ctx context.Context, username string) (*Users, error)
	FindByEmail(ctx context.Context, email string) (*Users, error)
	UpdatePassword(ctx context.Context, id int64, newPassword string) error
	GetUserByGoogleUserIDWithLock(ctx context.Context, ggID string) (*Users, error)
}
