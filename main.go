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

	categoryService = services.NewCategoryServiceImpl(categoryCollection, ctx)
	categoryController = controllers.NewCategoryController(categoryService)

	if err != nil {
		log.Fatal(err.Error())
	}

	server := gin.Default()

	basepath := server.Group("/v1")
	categoryController.RegisterCategoryRoutes(basepath)

	server.Run(":" + port)
}
