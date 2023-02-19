package main

import (
	"apiass/db"
	"apiass/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.Database.Close()
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, "hello") })

	// customer routes
	server.POST("/createCustomer", routes.NewCustomer)
	server.PUT("/updateCustomer/:id", routes.UpdateCustomer)
	server.GET("/viewCustomer/:id", routes.ViewCustomer)
	server.DELETE("/deleteCustomer/:id", routes.DeleteCustomer)

	// bank routes
	server.POST("/createBank", routes.NewBank)
	server.PUT("/updateBank/:id", routes.UpdateBank)
	server.GET("/viewBank/:id", routes.ViewBank)
	server.DELETE("/deleteBank/:id", routes.DeleteBank)

	// account routes
	server.POST("/openAccount", routes.NewAccount)
	server.POST("/openJointAccount", routes.NewJointAccount)
	server.GET("/viewAccount/:number", routes.ViewAccount)
	server.GET("/viewAccounttransactions/:number", routes.ViewAccounttransactions)
	server.DELETE("/deleteAccount/:number", routes.DeleteAccount)

	// transaction routes
	server.POST("/performTransaction", routes.PerformTransaction)
	server.GET("/viewTransaction/:id", routes.ViewTransaction)

	server.Run(":8080")

}
