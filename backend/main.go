package main

import (
	"fmt"
	"time"

	controller "github.com/AutoResto/controller"
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

	menu := router.Group("/menu")
	{
		// Melihat daftar menu
		menu.GET("/getmenu", controller.GetListMenu)
		// Pencarian menu
		menu.GET("/searchmenu", controller.SearchMenu)
		// Menambah menu baru
		menu.POST("/insertmenu", controller.InsertMenu)
		// Memperbarui menu
		menu.PUT("/updatemenu/:menu_id", controller.UpdateMenu)
		// Menghapus menu
		menu.DELETE("/deletemenu/:menu_id", controller.DeleteMenu)
	}
	recipe := router.Group("/recipe")
	{
		// Melihat resep menu
		recipe.GET("/getrecipe/:menu_id", controller.GetMenuRecipe)
	}

	material := router.Group("/material")
	{
		// Melihat bahan baku
		material.GET("/getmaterial/:material_id", controller.GetMaterial)
		// Pencarian bahan baku
		material.GET("/searchmaterial", controller.SearchMaterial)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
