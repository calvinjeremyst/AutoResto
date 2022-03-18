package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	cf "github.com/AutoResto/chef/controller/chef"
	iv "github.com/AutoResto/inventory/controller/inventory"
	ow "github.com/AutoResto/owner/controller/owner"
	entity "github.com/AutoResto/user/entity"
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
		InventoryManager.POST("/insert", iv.AddNewMaterial)
		InventoryManager.GET("/allmaterial", iv.ShowMaterial)
		InventoryManager.PUT("/:id", iv.UpdateMaterial)
		InventoryManager.DELETE("/:id", iv.RemoveMaterial)
	}

	//Chef Manager
	ChefManager := router.Group("/ChefManager")
	{
		ChefManager.GET("/allmenu", cf.ShowListMenu)
		ChefManager.GET("/:id", cf.ShowRecipeandMenu)
	}

	
	OwnerManager :=   router.Group("/OwnerManager")
	{
		OwnerManager.GET("/menu/:menu_name", ow.SearchMenu)
		OwnerManager.GET("/material/:material_name", ow.SearchMaterial)
		OwnerManager.GET("/allmaterial", iv.ShowMaterial)
		OwnerManager.GET("/allmenu", cf.ShowListMenu)
		OwnerManager.GET("/:id", ow.ShowAllRecipe)
		OwnerManager.POST("/insert", ow.AddNewMenu)
		OwnerManager.PUT("/:menu_id", ow.EditMenu)
		OwnerManager.DELETE("/:menu_id", ow.DeleteMenu)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
