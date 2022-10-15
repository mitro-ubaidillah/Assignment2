package main

import (
	"Tugas2/config"
	"Tugas2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/order/:id", inDB.UpdateOrder)
	router.DELETE("/order/:id", inDB.DeleteOrder)

	router.Run(":3000")
}
