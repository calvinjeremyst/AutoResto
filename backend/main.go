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
		InventoryManager.POST("/insert", controller.InsertMaterial)
		InventoryManager.GET("/allmaterial", controller.GetAllMaterial)
		InventoryManager.PUT("/:material_id", controller.UpdateMaterial)
		InventoryManager.DELETE("/:material_id", controller.DeleteMaterial)
	}

	//Chef Manager
	ChefManager := router.Group("/ChefManager")
	{
		ChefManager.GET("/allmenu", controller.GetListMenu)
		ChefManager.GET("/:menu_id", controller.GetMenuRecipe)
	}

	//Owner Manager
	OwnerManager := router.Group("/OwnerManager")
	{
		OwnerManager.GET("/menu/:menu_name", controller.SearchMenu)
		OwnerManager.GET("/material/:material_name", controller.SearchMaterial)
		OwnerManager.GET("/allmaterial", controller.GetAllMaterial)
		OwnerManager.GET("/allmenu", controller.GetListMenu)
		OwnerManager.GET("/:recipe_name", controller.GetAllRecipe)
		OwnerManager.POST("/insert", controller.InsertMenu)
		OwnerManager.PUT("/:menu_id", controller.UpdateMenu)
		OwnerManager.DELETE("/:menu_id", controller.DeleteMenu)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
