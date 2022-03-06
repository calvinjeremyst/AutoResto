package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	controller "github.com/AutoResto/controller"
	entity "github.com/AutoResto/module/user/entity"
)

func main() {
	fmt.Println("AutoResto Restaurant Management System")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/login", entity.Login)

	// Inventory Manager
	InventoryManager := router.Group("/InventoryManager")
	{
		InventoryManager.GET("/:name", controller.SearchMaterial)
		InventoryManager.POST("/insert", controller.InsertMaterial)
		InventoryManager.GET("/allmaterial", controller.GetAllMaterial)
		InventoryManager.PUT("/:material_id", controller.UpdateMaterial)
		InventoryManager.DELETE("/:material_id", controller.DeleteMaterial)
	}

	//Chef Manager
	ChefManager := router.Group("/ChefManager")
	{
		ChefManager.GET("")
	}

	//Owner Manager
	OwnerManager := router.Group("/OwnerManager")
	{
		OwnerManager.GET("/allmenu", controller.GetAllMenus)
		OwnerManager.POST("/insert", controller.InsertMenu)
		OwnerManager.DELETE("/:menu_id", controller.DeleteMenu)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
