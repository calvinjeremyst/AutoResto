package chef

import (
	"log"
	"net/http"

	modelMenu "github.com/AutoResto/menu/entity"
	mnservice "github.com/AutoResto/menu/service"
	modelRecipe "github.com/AutoResto/recipe/entity"
	rcservice "github.com/AutoResto/recipe/service"
	"github.com/gin-gonic/gin"
)

//Get List Menu
func ShowListMenu(c *gin.Context) {
	rows, err := mnservice.SelectMenuServiceDB()
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
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Get Menu Query Error"
		c.JSON(http.StatusBadRequest, response)
	}
}

func ShowRecipeandMenu(c *gin.Context) {

	rows, err := rcservice.SelectMenuRecipeServiceDB(c)
	if err != nil {
		log.Println(err)
	}

	var recipe modelRecipe.Recipe
	var recipes []modelRecipe.Recipe
	for rows.Next() {
		if err := rows.Scan(&recipe.Id, &recipe.Description, &recipe.Menu.Id, &recipe.Menu.Name, &recipe.Menu.Price); err != nil {
			log.Fatal(err)
		} else {
			recipes = append(recipes, recipe)
		}
	}
	var responseRecipe modelRecipe.RecipeResponse
	if err == nil {
		responseRecipe.Message = "Get Recipe Success"
		responseRecipe.Data = recipes
		c.JSON(http.StatusOK, responseRecipe)
	} else {
		responseRecipe.Message = "Recipe Get Failed"
		c.JSON(http.StatusBadRequest, responseRecipe)
	}
}
