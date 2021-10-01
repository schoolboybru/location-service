package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/location-service/db"
	adding "github.com/schoolboybru/location-service/internal/domain"
	"github.com/schoolboybru/location-service/internal/http/rest"
)

func main() {

	repository, err := db.New()

	if err != nil {
		panic(err)
	}

	service := adding.New(repository)

	handler := rest.NewHandler(service)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		locationGroup := v1.Group("/location")
		{
			locationGroup.GET("/:id", handler.Get)
			locationGroup.POST("/addLocation", handler.Post)
			locationGroup.GET("/country/:id", handler.GetByCountry)
			locationGroup.GET("city/:id", handler.GetByCity)
		}
	}

	router.Run(":8000")
}
