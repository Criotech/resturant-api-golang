package services

import (
	"github/criotech/resturant-api/types"
	"github/criotech/resturant-api/utils"
	"mime/multipart"
)

type FileUploadService interface {
	FileUpload(file multipart.File) (types.FileResponse, error)
}

type media struct{}

func NewFileUpload() FileUploadService {
	return &media{}
}

func (*media) FileUpload(file multipart.File) (types.FileResponse, error) {
	//upload
	uploadResponse, err := utils.ImageUploadHelper(file)
	if err != nil {
		return types.FileResponse{}, err
	}
	return uploadResponse, nil
}
