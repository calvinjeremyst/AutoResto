package main

import (
	"fmt"
	"time"

	
	entity "github.com/AutoResto/module/user/entity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	inventoryCt "github.com/AutoResto/controller/inventory"
	chefCt "github.com/AutoResto/controller/chef"
	ownerCt "github.com/AutoResto/controller/owner"
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

	inventory := router.Group("/InventoryManager")
	{
		inventory.GET("/:name",inventoryCt.SearchMaterial)
		inventory.POST("/insert",inventoryCt.InsertMaterial)
		inventory.PUT("/:id",inventoryCt.UpdateMaterial)
		inventory.DELETE("/:id",inventoryCt.DeleteMaterial)
	}

	chef := router.Group("/ChefManager")
	{
		chef.GET("/:id",chefCt.GetRecipeandMenu)
		
	}
	
	owner := router.Group("/OwnerManager")
	{
		owner.GET("/:id",ownerCt.GetAllRecipe)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
