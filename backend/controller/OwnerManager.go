package owner

import (

	"log"
	"net/http"
	controller  "github.com/AutoResto/controller"
	modelRecipe "github.com/AutoResto/module/recipe/entity"
	"github.com/gin-gonic/gin"
)

func GetAllRecipe(c *gin.Context){

	db := controller.Connect()
	defer db.Close()

	id := c.Param("id")
	query := "SELECT recipedetail.id,recipedetail.quantity,recipedetail.unit,recipe.id,recipe.description,material.name  FROM recipedetail JOIN recipe ON recipedetail.idRecipeFK = recipe.id JOIN material ON recipedetail.idMaterialFK = material.id WHERE recipedetail.id = '"+id+"'"
	rows,err := db.Query(query)

	if err != nil {
		log.Println(err)
	}

	var recipedetail modelRecipe.RecipeDetail
	var recipedetails []modelRecipe.RecipeDetail

	
	for rows.Next(){
		if err := rows.Scan(&recipedetail.Id,&recipedetail.Quantity,&recipedetail.Unit,&recipedetail.Recipe.Id,&recipedetail.Recipe.Description,&recipedetail.Material.Name);err != nil{
			log.Fatal(err)
		}else{
			recipedetails = append(recipedetails, recipedetail)
		}

	}

	var response modelRecipe.RecipeDetailResponse
	if err == nil{
		response.Message = "Get Recipe Success"
		response.Data = recipedetails
		sendRecipeSuccessResponse(c,response)

	}else{
		response.Message = "Get All Recipe Failed"
		sendRecipeErrorResponse(c,response)
	}
}

func sendRecipeSuccessResponse(c *gin.Context,response modelRecipe.RecipeDetailResponse){
	c.JSON(http.StatusOK, response)

}

func sendRecipeErrorResponse(c *gin.Context,response modelRecipe.RecipeDetailResponse){
	c.JSON(http.StatusBadRequest,response)
}
