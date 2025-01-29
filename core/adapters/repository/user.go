package repository

import (
	"context"
	"errors"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *adapters.Pgsql
}

func NewRepositoryUser(db *adapters.Pgsql) domain.RepositoryUser {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, db *gorm.DB, user *domain.Users) error {
	return db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "google_user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"user_name", "password", "phone_number", "email", "nick_name",
				"avatar", "updated_at",
			}),
		},
	).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *domain.Users) error {
	return r.db.DB().WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DB().WithContext(ctx).Delete(&domain.Users{}, id).Error
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*domain.Users, error) {
	var user domain.Users
	err := r.db.DB().WithContext(ctx).First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.Users, error) {
	var user domain.Users
	err := r.db.DB().WithContext(ctx).Where("user_name = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.Users, error) {
	var user domain.Users
	err := r.db.DB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, id int64, newPassword string) error {
	return r.db.DB().WithContext(ctx).Model(&domain.Users{}).Where("id = ?", id).Update("password", newPassword).Error
}

func (r *userRepository) GetUserByGoogleUserIDWithLock(ctx context.Context, ggID string) (*domain.Users, error) {
	var user domain.Users
	err := r.db.DB().WithContext(ctx).Where("google_user_id = ?", ggID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
