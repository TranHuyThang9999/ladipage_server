package services

import (
	"context"
	"ladipage_server/apis/entities"
	"ladipage_server/common/logger"
	customerrors "ladipage_server/core/custom_errors"
	"ladipage_server/core/domain"
)

type FileDescriptorsService struct {
	file   domain.RepositoryFileDescriptors
	logger *logger.Logger
}

func NewFileDescriptorsService(file domain.RepositoryFileDescriptors,
	logger *logger.Logger,
) *FileDescriptorsService {
	return &FileDescriptorsService{
		logger: logger,
		file:   file,
	}
}

func (u *FileDescriptorsService) DeleteFileById(ctx context.Context, userID, fileID int64) *customerrors.CustomError {
	return nil
}
func (u *FileDescriptorsService) AddListFileByObjectID(ctx context.Context, req *entities.CreateFilesRequest) *customerrors.CustomError {
	return nil
}
