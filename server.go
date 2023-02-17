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

	server.POST("/openAccount", routes.NewAccount)
	server.POST("/openJointAccount", routes.NewJointAccount)
	server.GET("/viewAccount/:number", routes.ViewAccount)
	server.GET("/viewAccountTransections/:number", routes.ViewAccountTransections)
	server.DELETE("/deleteAccount/:number", routes.DeleteAccount)

	server.Run(":8080")

}