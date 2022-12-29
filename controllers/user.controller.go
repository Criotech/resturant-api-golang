package controllers

import (
	"errors"
	middleware "github/criotech/resturant-api/middlewares"
	"github/criotech/resturant-api/models"
	"github/criotech/resturant-api/services"
	"github/criotech/resturant-api/types"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (pc *UserController) CreateAccount(ctx *gin.Context) {
	var user models.User

	user.ID = primitive.NewObjectID()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	validationErr := validate.Struct(&user)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userByEmail, _ := pc.UserService.GetUserByEmail(user.Email)

	if userByEmail != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, errors.New("user already exist"))
		ctx.JSON(http.StatusBadRequest, res)
	}

	hashPassword, err := utils.HashPassword(*user.Password)
	user.Password = &hashPassword

	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	result, err := pc.UserService.CreateAccount(&user)
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	res := utils.NewHTTPResponse(http.StatusCreated, result)

	ctx.JSON(http.StatusCreated, res)
}

func (pc *UserController) Login(ctx *gin.Context) {
	var loginCredentials types.LoginRequest

	if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	validationErr := validate.Struct(&loginCredentials)
	if validationErr != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := pc.UserService.GetUserByEmail(loginCredentials.Email)

	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, errors.New("auth failed"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	passwordIsValid := utils.VerifyPassword(*loginCredentials.Password, *user.Password)

	if !passwordIsValid {
		res := utils.NewHTTPResponse(http.StatusBadRequest, errors.New("auth failed"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token, refreshToken, err := utils.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.ID.String())
	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = pc.UserService.UpdateUserTokens(token, refreshToken, user.ID)
	user.Token = &token
	user.Refresh_token = &refreshToken

	if err != nil {
		res := utils.NewHTTPResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHTTPResponse(http.StatusOK, user)

	ctx.JSON(http.StatusOK, res)
}

func (pc *UserController) AuthRoute(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"success": "Access granted for api-2"})

}

func (userController *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	userroute.POST("/create", userController.CreateAccount)
	userroute.POST("/login", userController.Login)

	userroute.Use(middleware.Authenticate())
	userroute.GET("/auth", userController.AuthRoute)
}
