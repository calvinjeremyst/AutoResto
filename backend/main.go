package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	cf "github.com/AutoResto/controller/chef"
	iv "github.com/AutoResto/controller/inventory"
	ow "github.com/AutoResto/controller/owner"
	entity "github.com/AutoResto/domain/user/entity"
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
		InventoryManager.POST("/insert", iv.InsertMaterial)
		InventoryManager.GET("/allmaterial", iv.GetMaterial)
		InventoryManager.PUT("/:id", iv.UpdateMaterial)
		InventoryManager.DELETE("/:id", iv.DeleteMaterial)
	}

	//Chef Manager
	ChefManager := router.Group("/ChefManager")
	{
		ChefManager.GET("/allmenu", cf.GetListMenu)
		ChefManager.GET("/:id", cf.GetRecipeandMenu)
	}

	//Owner Manager
	OwnerManager := router.Group("/OwnerManager")
	{
		OwnerManager.GET("/menu/:menu_name", ow.SearchMenu)
		OwnerManager.GET("/material/:material_name", ow.SearchMaterial)
		OwnerManager.GET("/allmaterial", iv.GetMaterial)
		OwnerManager.GET("/allmenu", cf.GetListMenu)
		OwnerManager.GET("/:recipe_name", ow.GetAllRecipe)
		OwnerManager.POST("/insert", ow.InsertMenu)
		OwnerManager.PUT("/:menu_id", ow.UpdateMenu)
		OwnerManager.DELETE("/:menu_id", ow.DeleteMenu)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
