package repository

import (
	"context"
	"errors"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"

	"gorm.io/gorm"
)

type fileDescRepository struct {
	db *adapters.Pgsql
}

func NewRepositoryFileDesc(db *adapters.Pgsql) domain.RepositoryFileDescriptors {
	return &fileDescRepository{
		db: db,
	}
}

func (f *fileDescRepository) AddListFileWithTransaction(ctx context.Context, db *gorm.DB,
	files []*domain.FileDescriptors) error {
	result := db.WithContext(ctx).Create(files)
	return result.Error
}

func (f *fileDescRepository) AddWithTransaction(ctx context.Context, db *gorm.DB,
	file *domain.FileDescriptors) error {
	if err := db.WithContext(ctx).Create(file).Error; err != nil {
		return err
	}
	return nil
}

func (f *fileDescRepository) Add(ctx context.Context, file *domain.FileDescriptors) error {
	if err := f.db.DB().WithContext(ctx).Create(file).Error; err != nil {
		return err
	}
	return nil
}

func (f *fileDescRepository) DeleteFileByID(ctx context.Context, fileID, userID int64) error {
	if err := f.db.DB().WithContext(ctx).
		Where("id = ? and creator_id = ?", fileID, userID).
		Delete(&domain.FileDescriptors{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *fileDescRepository) DeleteFileByObjectID(ctx context.Context, objectID, userID int64) error {
	var file domain.FileDescriptors
	if err := f.db.DB().WithContext(ctx).
		Where("object_id = ? and creator_id = ?", objectID, userID).
		Delete(&file).Error; err != nil {
		return err
	}
	return nil
}

func (f *fileDescRepository) ListByObjectID(ctx context.Context, objectID int64) ([]*domain.FileDescriptors, error) {
	var files []*domain.FileDescriptors
	if err := f.db.DB().WithContext(ctx).
		Where("object_id = ?", objectID).
		Find(&files).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return files, nil
}

func (f *fileDescRepository) AddListFileWith(ctx context.Context, files []*domain.FileDescriptors) error {
	if err := f.db.DB().WithContext(ctx).Create(files).Error; err != nil {
		return err
	}

	return nil
}

func (f *fileDescRepository) DeleteListFileByObjectID(ctx context.Context, fileIds []int64) error {
	if err := f.db.DB().WithContext(ctx).
		Where("id in (?)", fileIds).
		Delete(&domain.FileDescriptors{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *fileDescRepository) ListByObjectIDs(ctx context.Context, objectIDs []int64) ([]*domain.FileDescriptors, error) {
	var files []*domain.FileDescriptors
	err := r.db.DB().WithContext(ctx).Where("object_id IN ?", objectIDs).Find(&files).Error
	return files, err
}
