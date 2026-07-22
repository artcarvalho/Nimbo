package main

import (
	"nimbo-api/controller"
	"nimbo-api/db"
	"nimbo-api/repository"
	"nimbo-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	CaptureRepository := repository.NewCaptureRepository(dbConnection)

	CaptureUsecase := useCase.NewCaptureUseCase(CaptureRepository)

	CaptureController := controller.NewCaptureController(CaptureUsecase)

	server.GET("/health", func(c *gin.Context) {
		c.String(200, "Server healthless")
	})

	server.GET("/capture", CaptureController.GetCapture)

	server.Run(":8080")
}
