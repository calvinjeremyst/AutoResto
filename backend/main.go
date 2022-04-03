package main

import (
	"fmt"
	"time"

	cf "github.com/AutoResto/chef/controller/chef"
	iv "github.com/AutoResto/inventory/controller/inventory"
	ow "github.com/AutoResto/owner/controller/owner"
	entity "github.com/AutoResto/user/entity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	router.POST("/login", entity.Login)


	// Inventory Manager
	InventoryManager := router.Group("/InventoryManager")
	{
		InventoryManager.POST("/insert", iv.AddNewMaterial)//done
		InventoryManager.GET("/allmaterial", iv.ShowMaterial)//done
		InventoryManager.GET("/:id",iv.ShowMaterialById)//done
		InventoryManager.PUT("/:id",iv.UpdateMaterial)
		InventoryManager.DELETE("/:id", iv.RemoveMaterial)//done
	}

	//Chef Manager
	ChefManager := router.Group("/ChefManager")
	{
		ChefManager.GET("/allmenu", cf.ShowListMenu)//done
		ChefManager.GET("/:id", cf.ShowRecipeandMenu)//done
	}

	OwnerManager := router.Group("/OwnerManager")
	{
		OwnerManager.GET("/menu/:menu_name", ow.SearchMenu)//done
		OwnerManager.GET("/material/:material_name", ow.SearchMaterial)//done
		OwnerManager.GET("/allmaterial", iv.ShowMaterial)//done
		OwnerManager.GET("/allmenu", cf.ShowListMenu)//done
		OwnerManager.GET("/:id", ow.ShowAllRecipe)//done
		OwnerManager.POST("/insert", ow.AddNewMenu)
		OwnerManager.PUT("/:menu_id", ow.EditMenu)
		OwnerManager.DELETE("/:menu_id", ow.DeleteMenu)//done
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
