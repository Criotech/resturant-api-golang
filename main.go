package main

import (
	"context"
	"github/criotech/resturant-api/controllers"
	"github/criotech/resturant-api/database"
	"github/criotech/resturant-api/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx                context.Context
	categoryService    services.CategoryService
	categoryCollection *mongo.Collection
	categoryController controllers.CategoryController
)

var (
	productService    services.ProductService
	productCollection *mongo.Collection
	productController controllers.ProductController
)

var (
	userService    services.UserService
	userCollection *mongo.Collection
	userController controllers.UserController
)

var (
	fileUploadService    services.FileUploadService
	fileUploadController controllers.FileUploadController
)

var (
	orderService    services.OrderService
	orderCollection *mongo.Collection
	orderController controllers.OrderController
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	//db connection
	mongoClient := database.DBinstance()
	categoryCollection = mongoClient.Database("resturant").Collection("categories")
	productCollection = mongoClient.Database("resturant").Collection("products")
	userCollection = mongoClient.Database("resturant").Collection("users")

	categoryService = services.NewCategoryServiceImpl(categoryCollection, ctx)
	categoryController = controllers.NewCategoryController(categoryService)

	productService = services.NewProductServiceImpl(productCollection, ctx)
	productController = controllers.NewProductController(productService)

	userService = services.NewUserServiceImpl(userCollection, ctx)
	userController = controllers.NewUserController(userService)

	fileUploadService = services.NewFileUpload()
	fileUploadController = controllers.NewFileUploadController(fileUploadService)

	orderService = services.NewOrderServiceImpl(orderCollection, ctx)
	orderController = controllers.NewOrderController(orderService)

	if err != nil {
		log.Fatal(err.Error())
	}

	server := gin.Default()

	basepath := server.Group("/v1")
	categoryController.RegisterCategoryRoutes(basepath)
	productController.RegisterProdutRoutes(basepath)
	userController.RegisterUserRoutes(basepath)
	fileUploadController.RegisterFileUploadRoutes(basepath)
	orderController.RegisterOrderRoutes(basepath)

	server.Run(":" + port)
}
