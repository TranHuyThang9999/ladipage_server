package controllers

import (
	"ladipage_server/apis/resources"
	"ladipage_server/core/services"
)

type FileDescController struct {
	file *services.FileDescriptorsService
	base *baseController
	reso *resources.Resource
}

func NewFileDescController(file *services.FileDescriptorsService,
	base *baseController,
	reso *resources.Resource) *FileDescController {
	return &FileDescController{
		file: file,
	}
}
