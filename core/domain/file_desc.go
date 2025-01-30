package domain

import (
	"context"
	"ladipage_server/apis/entities"

	"gorm.io/gorm"
)

type FileDescriptors struct {
	*entities.Model
	CreatorID  int64  `json:"creator_id,omitempty"`
	ObjectID   int64  `json:"object_id,omitempty"`
	Url        string `json:"url,omitempty"`
	TypeObject int8   `json:"type,omitempty"`
}

func (FileDescriptors) TableName() string {
	return "file_descriptors"
}

type RepositoryFileDescriptors interface {
	Add(ctx context.Context, file *FileDescriptors) error
	ListByObjectID(ctx context.Context, objectID int64) ([]*FileDescriptors, error)
	DeleteFileByID(ctx context.Context, fileID, userID int64) error
	DeleteFileByObjectID(ctx context.Context, objectID, userID int64) error
	AddWithTransaction(ctx context.Context, db *gorm.DB, file *FileDescriptors) error
	AddListFileWithTransaction(ctx context.Context, db *gorm.DB, files []*FileDescriptors) error
	AddListFileWith(ctx context.Context, files []*FileDescriptors) error
	DeleteListFileByObjectID(ctx context.Context, fileIds []int64) error
}
