package controllers

import (
	"github/criotech/resturant-api/services"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileUploadController struct {
	FileService services.FileUploadService
}

func NewFileUploadController(fileservice services.FileUploadService) FileUploadController {
	return FileUploadController{
		FileService: fileservice,
	}
}

func (c *FileUploadController) UploadFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	src, err := file.Open()
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	defer src.Close()
	fileResponse, err := c.FileService.FileUpload(src)

	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHTTPResponse(http.StatusOK, fileResponse)

	ctx.JSON(http.StatusOK, res)
}

func (c *FileUploadController) RegisterFileUploadRoutes(rg *gin.RouterGroup) {
	fileuploadroute := rg.Group("/fileupload")
	fileuploadroute.POST("/", c.UploadFile)
}
