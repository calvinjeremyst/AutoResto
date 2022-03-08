package controller

import (
	"log"

	modelMenu "github.com/AutoResto/module/menu/entity"
	modelRecipe "github.com/AutoResto/module/recipe/entity"

	"github.com/gin-gonic/gin"
)

//Get List Menu
func GetListMenu(c *gin.Context) {
	db := Connect()
	defer db.Close()

	query := "SELECT id,name,price FROM menu"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var menu modelMenu.Menu
	var menus []modelMenu.Menu

	for rows.Next() {
		if err := rows.Scan(&menu.Id, &menu.Name, &menu.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			menus = append(menus, menu)
		}
	}

	var response modelMenu.MenuResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = menus
		sendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendMenuErrorResponse(c, response)
	}
}

// Get Menu Recipe
func GetMenuRecipe(c *gin.Context) {
	db := Connect()
	defer db.Close()

	menuId := c.Param("menu_id")

	query := "SELECT recipedetail.id, recipedetail.quantity, recipedetail.idMaterialFK, recipedetail.idRecipeFK, recipedetail.unit, material.name FROM `recipedetail` join `material` on recipedetail.id = material.id WHERE recipedetail.idRecipeFK=" + menuId + ""

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var recipedetail modelRecipe.RecipeDetail
	var recipedetails []modelRecipe.RecipeDetail

	for rows.Next() {
		if err := rows.Scan(&recipedetail.Id, &recipedetail.Quantity, &recipedetail.Material.Id, &recipedetail.Recipe.Id, &recipedetail.Unit, &recipedetail.Material.Name); err != nil {
			log.Fatal(err.Error())
		} else {
			recipedetails = append(recipedetails, recipedetail)
		}
	}

	var response modelRecipe.RecipeDetailResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = recipedetails
		sendRecipeDetailSuccessResponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendRecipeDetailErrorResponse(c, response)
	}
}
