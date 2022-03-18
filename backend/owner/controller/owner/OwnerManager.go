package owner

import (
	"log"
	"net/http"
	"strings"
	modelMaterial "github.com/AutoResto/material/entity"
	modelRecipe "github.com/AutoResto/recipe/entity"
	mtservice "github.com/AutoResto/material/service"
	modelMenu "github.com/AutoResto/menu/entity"
	mnservice "github.com/AutoResto/menu/service"
	rcservice "github.com/AutoResto/recipe/service"
	"github.com/gin-gonic/gin"
)

// Search Menu
func SearchMenu(c *gin.Context) {
	rows,err := mnservice.SearchMenuServiceDB(c)
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
		c.JSON(http.StatusOK,response)
	} else {
		response.Message = "Get Menu Query Error"
		c.JSON(http.StatusBadRequest,response)
	}
}

// Search Material
func SearchMaterial(c *gin.Context) {
	rows,err := mtservice.SearchMaterialServiceDB(c)
	if err != nil {
		log.Println(err)
	}
	var material modelMaterial.Material
	var materials []modelMaterial.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}
	var response modelMaterial.MaterialResponse
	if err == nil {
		response.Message = "Get Material Success"
		response.Data = materials
		c.JSON(http.StatusOK,response)
		
	} else {
		response.Message = "Get Material Query Error"
		c.JSON(http.StatusBadRequest,response)
	}
}
//Get All Recipe
func ShowAllRecipe(c *gin.Context) {
	rows,err := rcservice.SelectAllRecipeServiceDB(c)
	if err != nil {
		log.Println(err)
	}

	var recipe modelRecipe.RecipeDetail
	var recipes []modelRecipe.RecipeDetail

	for rows.Next() {
		if err := rows.Scan(&recipe.Id, &recipe.Quantity, &recipe.Unit, &recipe.Recipe.Id,&recipe.Recipe.Description,&recipe.Material.Name); err != nil {
			log.Fatal(err.Error())
		} else {
			recipes = append(recipes, recipe)
		}
	}

	var response modelRecipe.RecipeDetailResponse
	if err == nil {
		response.Message = "Get Material Success"
		response.Data = recipes
		c.JSON(http.StatusOK,response)
	} else {
		response.Message = "Get Material Query Error"
		c.JSON(http.StatusBadRequest,response)
	}
}

// Insert Menu
func AddNewMenu(c *gin.Context) {
	materialName := c.PostForm("material")
	materialArr := strings.Split(materialName, ",")
	
	//insert recipe
	errQuery := rcservice.InsertRecipeService(c)

	//insert menu dengan select id dari description
	errQuery = mnservice.InsertMenuService(c)

	//mencari material dengan nama yang sama
	rows,err := mtservice.SearchMaterialServiceDB(c)
	if err != nil {
		log.Println(err)
	}

	var material modelMaterial.Material
	var materials []modelMaterial.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}

	for i := 0; i < len(materialArr); i++ {
		for j := 0; j < len(materials); j++ {
			//mengecek apakah material sudah ada
			if materialArr[i] != materials[j].Name {
				// menambah material jika tidak ada
				errQuery = mtservice.InsertMaterialHelperServiceDB(c,i)
				
			}
		}
		/*_, errQuery = db.Exec("INSERT INTO recipedetail(quantity,idMaterialFK,idRecipeFK,unit) VALUES(?,(SELECT id FROM material WHERE material.name=?),(SELECT id FROM recipe WHERE recipe.description=?),?)",
			quantityArr[i],
			materialArr[i],
			description,
			unitArr[i],
		)*/
		errQuery = rcservice.InsertRecipeDetailService(c,i)
	}

	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Insert Menu Success"
		c.JSON(http.StatusOK,response)
	} else {
		response.Message = "Insert Menu Failed Error"
		c.JSON(http.StatusBadRequest,response)
	}
}

//Update Menu
func EditMenu(c *gin.Context) {
	
	errQuery := mnservice.UpdateMenuServiceDB(c)

	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Update Menu Success"
		c.JSON(http.StatusOK,response)
	} else {
		response.Message = "Update Menu Failed Error"
		c.JSON(http.StatusBadRequest,response)
	}
}

// Delete Menu
func DeleteMenu(c *gin.Context) {

	errQuery := mnservice.DeleteMenuServiceDB(c)
	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Delete Menu Success"
		c.JSON(http.StatusOK,response)
	} else {
		response.Message = "Delete Menu Failed Error"
		c.JSON(http.StatusBadRequest,response)
	}
}
