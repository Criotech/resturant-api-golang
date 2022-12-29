package controllers

import (
	middleware "github/criotech/resturant-api/middlewares"
	"github/criotech/resturant-api/models"
	"github/criotech/resturant-api/services"
	"github/criotech/resturant-api/types"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productservice services.ProductService) ProductController {
	return ProductController{
		ProductService: productservice,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	product.ID = primitive.NewObjectID()

	if err := ctx.ShouldBindJSON(&product); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	validationErr := validate.Struct(&product)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := pc.ProductService.CreateProduct(&product)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	res := utils.NewHTTPResponse(http.StatusCreated, result)

	ctx.JSON(http.StatusCreated, res)
}

func (uc *ProductController) GetProduct(ctx *gin.Context) {
	var productID string = ctx.Param("productId")

	err := utils.CheckUserType(ctx, "ADMIN")
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	product, err := uc.ProductService.GetProduct(&productID)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHTTPResponse(http.StatusOK, product)
	ctx.JSON(http.StatusOK, res)
}

func (uc *ProductController) GetProducts(ctx *gin.Context) {
	var productQueries types.QueryProducts

	if err := ctx.ShouldBindQuery(&productQueries); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	products, err := uc.ProductService.GetProducts(productQueries)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHTTPResponse(http.StatusOK, products)
	ctx.JSON(http.StatusOK, res)
}

func (uc *ProductController) UpdateProduct(ctx *gin.Context) {
	var req types.UpdateProductRequest
	var productID string = ctx.Param("productId")

	err := utils.CheckUserType(ctx, "ADMIN")
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	validationErr := validate.Struct(&req)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.ProductService.UpdateProduct(&productID, &req)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHTTPResponse(http.StatusOK, "successful")
	ctx.JSON(http.StatusOK, res)
}

func (uc *ProductController) DeleteProduct(ctx *gin.Context) {
	var productID string = ctx.Param("productId")
	err := utils.CheckUserType(ctx, "ADMIN")
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.ProductService.DeleteProduct(&productID)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHTTPResponse(http.StatusOK, "successful")
	ctx.JSON(http.StatusOK, res)
}

func (productController *ProductController) RegisterProdutRoutes(rg *gin.RouterGroup) {
	productroute := rg.Group("/products")
	productroute.GET("/", productController.GetProducts)
	productroute.GET("/:productId", productController.GetProduct)

	productroute.Use(middleware.Authenticate())
	productroute.POST("/", productController.CreateProduct)
	productroute.PUT("/:productId", productController.UpdateProduct)
	productroute.DELETE("/:productId", productController.DeleteProduct)
}
