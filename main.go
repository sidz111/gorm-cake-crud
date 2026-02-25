package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/cake-gorm/config"
	"github.com/sidz111/cake-gorm/controller"
	"github.com/sidz111/cake-gorm/model"
	"github.com/sidz111/cake-gorm/repository"
	"github.com/sidz111/cake-gorm/service"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect db")
	}
	config.DB.AutoMigrate(&model.Cake{})

	repo := repository.NewCakeRepository(config.DB)
	serv := service.NewCakeService(repo)
	cont := controller.NewCakeController(serv)

	r := gin.Default()

	cake := r.Group("/cakes")
	{
		cake.POST("/", cont.Create)
		cake.GET("/:id", cont.GetById)
		cake.PUT("/", cont.Update)
		cake.GET("/", cont.GetAll)
		cake.DELETE("/:id", cont.Delete)
	}
	r.Run(":8080")
}
