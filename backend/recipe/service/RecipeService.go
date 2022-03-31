package service

import (
	"database/sql"
	//"strings"
	rcrepo "github.com/AutoResto/recipe/repository"
	handler "github.com/AutoResto/handler"
	"github.com/gin-gonic/gin"
)

type RecipeRepo struct{
	rcrepo.RecipeRepository
}

func NewRecipeRepository() *RecipeRepo{
	return &RecipeRepo{}
}

func (r *RecipeRepo) InsertRecipeService(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()

	description := c.PostForm("description")

	_, errQuery := db.Exec("INSERT INTO recipe(description) VALUES (?)",
		description,
	)

	return errQuery
}

func(r *RecipeRepo)InsertRecipeDetailService(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()

	
	quantity := c.PostForm("quantity_recipe")

	//id_material := c.Param("id_material")
	materialName := c.PostForm("material")

	//id_recipe := c.Param("id_recipe")
	description := c.PostForm("description")
	unit := c.PostForm("unit_recipe")

	//materialArr := strings.Split(materialName, ",")
	//quantityArr := strings.Split(quantity, ",")
	//unitArr := strings.Split(unit, ",")

	_, errQuery := db.Exec("INSERT INTO recipedetail(quantity,idMaterialFK,idRecipeFK,unit) VALUES(?,(SELECT id FROM material WHERE material.name= '"+materialName+"'),(SELECT id FROM recipe WHERE recipe.description= '"+description+"'),?)",
		quantity,
		unit,
	)

	return errQuery
}

func (r *RecipeRepo) SelectMenuRecipeServiceDB(c *gin.Context) (row *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()

	id := c.Param("id")
	query := "SELECT recipe.id,recipe.description,menu.id,menu.name,menu.price FROM recipe INNER JOIN menu ON recipe.id = menu.idRecipeFK WHERE recipe.id = '" + id + "'"

	recipe, err := db.Query(query)
	return recipe, err
}

func(r *RecipeRepo) SelectAllRecipeServiceDB(c *gin.Context) (row *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()
	id := c.Param("id")
	query := "SELECT recipedetail.id,recipedetail.quantity,recipedetail.unit,recipe.id,recipe.description,material.name  FROM recipedetail JOIN recipe ON recipedetail.idRecipeFK = recipe.id JOIN material ON recipedetail.idMaterialFK = material.id WHERE recipedetail.id = '" + id + "'"
	recipe, err := db.Query(query)
	return recipe, err

}
