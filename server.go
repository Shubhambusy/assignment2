package main

import (
	"apiass/routes"
	"apiass/db"	

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.Database.Close()
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {   ctx.JSON(200, "hello")  })

	server.POST("/createCustomer", routes.NewCustomer)
	server.PUT("/updateCustomer/:id", routes.UpdateCustomer)
	server.GET("/viewCustomer/:id", routes.ViewCustomer)
	server.DELETE("/deleteCustomer/:id", routes.DeleteCustomer)

	server.POST("/createBank", routes.NewBank)
	server.PUT("/updateBank/:id", routes.UpdateBank)
	server.GET("/viewBank/:id", routes.ViewBank)
	server.DELETE("/deleteBank/:id", routes.DeleteBank)

	server.Run(":8080")

}