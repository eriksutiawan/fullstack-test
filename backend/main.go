package main

import (
	"backend/internal/app/submission/delivery"
	submissionusecase "backend/internal/app/submission/usecase"
	userrepository "backend/internal/app/user/repository"
	userusecase "backend/internal/app/user/usecase"
	"backend/internal/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	dbMongo := config.ConnectToMongo()

	userRepo := userrepository.NewMongoUserRepository(dbMongo)
	userUsecase := userusecase.NewUserService(userRepo)
	submissionUsecase := submissionusecase.NewSubmissionService(userUsecase)

	router := gin.Default()
	router.POST("/api/form", func(c *gin.Context) {
		delivery.SubmissionRegisterHandler(c, submissionUsecase)
	})
	router.Run(":8080")
}
