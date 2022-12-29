package types

import "mime/multipart"

type FileRequest struct {
	File multipart.FileHeader `form:"file,omitempty" validate:"required"`
}

type FileResponse struct {
	FileURL  string
	PublicId string
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}
