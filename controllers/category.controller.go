package controllers

import (
	"github/criotech/resturant-api/models"
	"github/criotech/resturant-api/services"
	"github/criotech/resturant-api/types"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryservice services.CategoryService) CategoryController {
	return CategoryController{
		CategoryService: categoryservice,
	}
}

var validate = validator.New()

func (uc *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category

	category.ID = primitive.NewObjectID()

	if err := ctx.ShouldBindJSON(&category); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	validationErr := validate.Struct(&category)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.CategoryService.CreateCategory(&category)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	res := utils.NewHTTPResponse(http.StatusOK, result)

	ctx.JSON(http.StatusCreated, res)
}

func (uc *CategoryController) GetCategory(ctx *gin.Context) {
	var categoryID string = ctx.Param("categoryId")
	category, err := uc.CategoryService.GetCategory(&categoryID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (uc *CategoryController) GetCategories(ctx *gin.Context) {
	categories, err := uc.CategoryService.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (uc *CategoryController) UpdateCategory(ctx *gin.Context) {
	var req types.UpdateCategoryRequest
	var categoryID string = ctx.Param("categoryId")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	validationErr := validate.Struct(&req)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	err := uc.CategoryService.UpdateCategory(&categoryID, &req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "success"})
}

func (uc *CategoryController) DeleteCategory(ctx *gin.Context) {
	var categoryID string = ctx.Param("categoryId")
	err := uc.CategoryService.DeleteCategory(&categoryID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (categoryController *CategoryController) RegisterCategoryRoutes(rg *gin.RouterGroup) {
	categoryroute := rg.Group("/categories")
	categoryroute.POST("/", categoryController.CreateCategory)
	categoryroute.GET("/", categoryController.GetCategories)
	categoryroute.GET("/:categoryId", categoryController.GetCategory)
	categoryroute.PUT("/:categoryId", categoryController.UpdateCategory)
	categoryroute.DELETE("/:categoryId", categoryController.DeleteCategory)
}
