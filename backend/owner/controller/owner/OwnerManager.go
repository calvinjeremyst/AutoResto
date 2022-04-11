package owner

import (
	"log"
	"net/http"

	modelMaterial "github.com/AutoResto/material/entity"
	mtservice "github.com/AutoResto/material/service"
	modelMenu "github.com/AutoResto/menu/entity"
	mnservice "github.com/AutoResto/menu/service"
	modelRecipe "github.com/AutoResto/recipe/entity"
	rcservice "github.com/AutoResto/recipe/service"
	"github.com/gin-gonic/gin"
)

// Search Menu
func SearchMenu(c *gin.Context) {
	rows, err := mnservice.NewMenuRepository().SearchMenuServiceDB(c)
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

// Search Material
func SearchMaterial(c *gin.Context) {
	rows, err := mtservice.NewMaterialRepository().SearchMaterialServiceDB(c)
	if err != nil {
		log.Println(err)
	}
	var material modelMaterial.Material
	var materials []modelMaterial.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit, &material.Inventory.Id); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}
	var response modelMaterial.MaterialResponse
	if err == nil {
		response.Message = "Get Material Success"
		response.Data = materials
		c.JSON(http.StatusOK, response)

	} else {
		response.Message = "Get Material Query Error"
		c.JSON(http.StatusBadRequest, response)
	}
}

// Insert Menu
func AddNewMenu(c *gin.Context) {
	materialName := c.PostForm("material")
	//materialArr := strings.Split(materialName, ",")
	errQuery := rcservice.NewRecipeRepository().InsertRecipeService(c)
	//insert recipe

	//insert menu dengan select id dari description
	errQuery = mnservice.NewMenuRepository().InsertMenuService(c)

	//mencari material dengan nama yang sama
	rows, err := mtservice.NewMaterialRepository().SearchMaterialServiceDB(c)
	if err != nil {
		log.Println(err)
	}

	var material modelMaterial.Material
	var materials []modelMaterial.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit, &material.Inventory.Id); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}
	cek := true
	for i := 0; i < len(materials); i++ {
		if materials[i].Name == materialName {
			cek = false
		}
	}

	if cek == true {
		errQuery = mtservice.NewMaterialRepository().InsertMaterialHelperServiceDB(c)
		errQuery = rcservice.NewRecipeRepository().InsertRecipeDetailService(c)
	}
	//for i := 0; i < len(materialArr); i++ {
	//	for j := 0; j < len(materials); j++ {
	//		//mengecek apakah material sudah ada
	//		if materialArr[i] != materials[j].Name {
	//			// menambah material jika tidak ada
	//			errQuery = mtservice.InsertMaterialHelperServiceDB(c, i)
	//			break
	//		} else {
	//			break
	//		}
	//	}
	//	errQuery = rcservice.InsertRecipeDetailService(c, i)
	//}*/

	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Insert Menu Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Insert Menu Failed Error"
		c.JSON(http.StatusBadRequest, response)
	}
}

func AddnewMenus(c *gin.Context) {

	rows, err := rcservice.NewRecipeRepository().GetAvailabilityMenuDescription(c)
	menuName := c.PostForm("name")
	description := c.PostForm("description")

	if err != nil {
		log.Println(err)
	}

	var recipe modelRecipe.Recipe
	var recipes []modelRecipe.Recipe

	for rows.Next() {
		if err := rows.Scan(&recipe.Menu.Id, &recipe.Menu.Name, &recipe.Menu.Price, &recipe.Description); err != nil {
			log.Fatal(err.Error())
		} else {
			recipes = append(recipes, recipe)
		}

	}

	cek := true
	for i := 0; i < len(recipes); i++ {
		if recipes[i].Description == description && recipes[i].Menu.Name == menuName {
			cek = false
		}

	}
	var response modelRecipe.RecipeResponse
	if cek == true {
		errQuery := rcservice.NewRecipeRepository().InsertRecipeService(c)
		errQuery2 := mnservice.NewMenuRepository().InsertMenuService(c)

		if errQuery == nil && errQuery2 == nil {
			response.Message = "Insert Menu And Description Success"
			c.JSON(http.StatusOK, response)
		} else {
			response.Message = "Insert Menu And Description Failed"
			c.JSON(http.StatusBadRequest, response)
		}
	}
}

//Update Menu
func EditMenu(c *gin.Context) {
	errQuery := mnservice.NewMenuRepository().UpdateMenuServiceDB(c)
	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Update Menu Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Update Menu Failed Error"
		c.JSON(http.StatusBadRequest, response)
	}
}

// Delete Menu
func DeleteMenu(c *gin.Context) {
	errQuery := mnservice.NewMenuRepository().DeleteMenuServiceDB(c)
	var response modelMenu.MenuResponse
	if errQuery == nil {
		response.Message = "Delete Menu Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Delete Menu Failed Error"
		c.JSON(http.StatusBadRequest, response)
	}
}
