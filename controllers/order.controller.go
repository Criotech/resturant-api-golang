package controllers

import (
	"encoding/json"
	"github/criotech/resturant-api/services"
	"github/criotech/resturant-api/types"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService services.OrderService
}

func NewOrderController(orderservice services.OrderService) OrderController {
	return OrderController{
		OrderService: orderservice,
	}
}

func (oc *OrderController) InitializePayment(ctx *gin.Context) {
	var pay types.InitializePayment

	if err := ctx.ShouldBindJSON(&pay); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	validationErr := validate.Struct(&pay)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := oc.OrderService.InitializePayment("opesiyanbola8991@gmail.com", pay.Amount)

	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		panic(err)
	}

	res := utils.NewHTTPResponse(http.StatusCreated, data["data"])

	ctx.JSON(http.StatusCreated, res)
}

func (orderController *OrderController) RegisterOrderRoutes(rg *gin.RouterGroup) {
	orderroute := rg.Group("/orders")
	orderroute.POST("/", orderController.InitializePayment)
}
