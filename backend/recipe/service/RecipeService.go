package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
	handler "github.com/AutoResto/handler"
)


func InsertRecipeService(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()

	description := c.PostForm("description")

	_, errQuery := db.Exec("INSERT INTO recipe(description) VALUES (?)",
		description,

	)

	return errQuery
}

func InsertRecipeDetailService(c *gin.Context,i int) error{
	db := handler.Connect()
	defer db.Close()

	materialName := c.PostForm("material")
	quantity := c.PostForm("quantity")
	description := c.PostForm("description")
	unit := c.PostForm("unit")

	materialArr := strings.Split(materialName, ",")
	quantityArr := strings.Split(quantity, ",")
	unitArr := strings.Split(unit, ",")

	_, errQuery := db.Exec("INSERT INTO recipedetail(quantity,idMaterialFK,idRecipeFK,unit) VALUES(?,(SELECT id FROM material WHERE material.name=?),(SELECT id FROM recipe WHERE recipe.description=?),?)",
			quantityArr[i],
			materialArr[i],
			description[i],
			unitArr[i],
		)

	return errQuery
}

func SelectMenuRecipeServiceDB(c* gin.Context) (row *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()
	id := c.Param("id")
	query := "SELECT recipe.id,recipe.description,menu.id,menu.name,menu.price FROM recipe INNER JOIN menu ON recipe.id = menu.idRecipeFK WHERE recipe.id = '" + id + "'"
	recipe, err := db.Query(query)
	return recipe,err
}

func SelectAllRecipeServiceDB(c *gin.Context)(row *sql.Rows,err error){
	db := handler.Connect()
	defer db.Close()
	id := c.Param("id")
	query := "SELECT recipedetail.id,recipedetail.quantity,recipedetail.unit,recipe.id,recipe.description,material.name  FROM recipedetail JOIN recipe ON recipedetail.idRecipeFK = recipe.id JOIN material ON recipedetail.idMaterialFK = material.id WHERE recipedetail.id = '"+id+"'"
	recipe,err := db.Query(query)
	return recipe,err

}