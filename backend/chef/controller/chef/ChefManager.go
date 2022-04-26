package chef

import (
	"log"
	"net/http"

	//modelMenu "github.com/AutoResto/menu/entity"
	mnservice "github.com/AutoResto/menu/service"
	modelRecipe "github.com/AutoResto/recipe/entity"
	rcservice "github.com/AutoResto/recipe/service"
	rcmodel "github.com/AutoResto/recipe/entity"
	mtservice "github.com/AutoResto/material/service"
	mtmodel "github.com/AutoResto/material/entity"
	"github.com/gin-gonic/gin"
)

//Get All Recipe
func ShowAllRecipe(c *gin.Context) {
	rows, err := rcservice.NewRecipeRepository().SelectAllRecipeServiceDB(c)
	if err != nil {
		log.Println(err)
	}

	var recipe modelRecipe.RecipeDetail
	var recipes []modelRecipe.RecipeDetail

	for rows.Next() {
		if err := rows.Scan(&recipe.Id, &recipe.Quantity, &recipe.Unit, &recipe.Recipe.Id, &recipe.Recipe.Description, &recipe.Material.Id, &recipe.Material.Name); err != nil {
			log.Fatal(err.Error())
		} else {
			recipes = append(recipes, recipe)
		}
	}

	var response modelRecipe.RecipeDetailResponse
	if err == nil {
		response.Message = "Get Recipe Detail Success"
		response.Data = recipes
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Get Recipe Detail Failed"
		c.JSON(http.StatusBadRequest, response)
	}
}

func ShowDescription(c *gin.Context){
	rows,err := rcservice.GetDescriptionService(c)
	if err != nil{
		log.Println(err)
	}

	var descname modelRecipe.Recipe
	var descnames []modelRecipe.Recipe

	for rows.Next(){
		if err := rows.Scan(&descname.Id,&descname.Description);err != nil{
			log.Fatal(err.Error)
		}else{
			descnames = append(descnames, descname)
		}
	}

	var response rcmodel.RecipeResponse
	if err == nil{
		response.Message = "Get Description Success"
		response.Data = descnames
		c.JSON(http.StatusOK,response)
	}else{
		response.Message = "Get Description Failed"
		c.JSON(http.StatusBadRequest,response)
	}
}

func ShowMaterialName(c *gin.Context){
	rows,err := mtservice.NewMaterialRepository().GetNameMaterialService()
	if err != nil{
		log.Println(err)
	}

	var material mtmodel.Material
	var listmaterial []mtmodel.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id,&material.Name);err !=nil{
			log.Fatal(err)
		}else{
			listmaterial = append(listmaterial, material)
		}
	}


	var response mtmodel.MaterialResponse
	if err == nil{
		response.Message = "Get material name success"
		response.Data = listmaterial
		c.JSON(http.StatusOK,response)
	}else{
		response.Message = "Get Material Failed"
		c.JSON(http.StatusBadRequest,response)
	}
}

func ShowMaterialNameAndDescription(c *gin.Context){
	rows,err := rcservice.SelectDescAndMaterialNameService(c)
	if err != nil{
		log.Println(err)
	}

	var object modelRecipe.RecipeDetail
	var objects []modelRecipe.RecipeDetail

	for rows.Next(){
		if err := rows.Scan(&object.Id,&object.Recipe.Id,&object.Recipe.Description,&object.Material.Id,&object.Material.Name);err != nil{
			log.Fatal(err)
		}else{
			objects = append(objects, object)
		}
	}

	var response modelRecipe.RecipeDetailResponse
	if err == nil{
		response.Message = "Get Description and Material Name success"
		response.Data = objects
		c.JSON(http.StatusOK,response)
	}else{
		response.Message = "Get Description and Material Name Failed"
		c.JSON(http.StatusBadRequest,response)
	}

}

func AddRecipeForEachMaterial(c *gin.Context) {
	var response modelRecipe.RecipeDetailResponse
	errQuery := rcservice.NewRecipeRepository().InsertRecipeDetailService(c)

	if errQuery == nil {
		response.Message = "Insert Recipe dan Bahan bahan Sukses"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Insert Recipe dan Bahan bahan Failed"
		c.JSON(http.StatusBadRequest, response)
	}
}

func AddRecipeandMaterial(c *gin.Context){
	
	var recipedt rcmodel.RecipeDetail
	var response modelRecipe.RecipeDetailResponse
	
	err := c.BindJSON(&recipedt)

	if err == nil {
		errQuery := rcservice.InsertDetailRecipeJSON(recipedt,c)
		if errQuery == nil{
			response.Message = "Insert Recipe dan Bahan bahan sukes"
			c.JSON(http.StatusOK,response)
		}else{
			response.Message = "Insert Recipe dan bahan bahan failed"
			c.JSON(http.StatusBadRequest,response)
		}
	}else{
		response.Message = "Failed Bind JSON"
		c.JSON(http.StatusBadRequest,response)
	}
}

//Get List Menu Owner
func ShowListMenuChef(c *gin.Context) {
	rows, err := mnservice.NewMenuRepository().SelectMenuChefServiceDB()
	if err != nil {
		log.Println(err)
	}
	var detailmenu modelRecipe.Recipe
	var detailmenus []modelRecipe.Recipe
	for rows.Next() {
		if err := rows.Scan(&detailmenu.Menu.Id, &detailmenu.Menu.Name,&detailmenu.Description); err != nil {
			log.Fatal(err.Error())
		} else {
			detailmenus = append(detailmenus, detailmenu)
		}
	}
	var response modelRecipe.RecipeResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = detailmenus
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Get Menu Query Error"
		c.JSON(http.StatusBadRequest, response)
	}
}

func ShowRecipeandMenu(c *gin.Context) {

	rows, err := rcservice.NewRecipeRepository().SelectMenuRecipeServiceDB(c)
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
