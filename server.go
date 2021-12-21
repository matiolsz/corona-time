package main

import (
	"corona-time/controller"

	"github.com/gin-gonic/gin"
)

var (
	covidController controller.CovidController = controller.New()
)

func main() {

	server := gin.New()

	server.Static("/css", ".templates/css")

	server.LoadHTMLGlob("templates/*.html")

	viewForToday := server.Group("/view")
	{
		viewForToday.GET("/usa", covidController.ShowForToday)
	}
	viewHistory := server.Group("/view")
	{
		viewHistory.GET("/history", covidController.ShowHistory)
	}

	server.Run(":8080")

}
