package utils

import (
	"context"
	"github/criotech/resturant-api/config"
	"github/criotech/resturant-api/types"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input interface{}) (types.FileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return types.FileResponse{}, err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return types.FileResponse{}, err
	}
	return types.FileResponse{FileURL: uploadParam.SecureURL, PublicId: uploadParam.PublicID}, nil
}
