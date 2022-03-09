package chef

import (
	"log"
	"strconv"

	"github.com/AutoResto/controller"
	modelMenu "github.com/AutoResto/module/menu/entity"
	modelRecipe "github.com/AutoResto/module/recipe/entity"
	"github.com/gin-gonic/gin"
)

//Get List Menu
func GetListMenu(c *gin.Context) {
	db := controller.Connect()
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
		controller.SendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		controller.SendMenuErrorResponse(c, response)
	}
}

func GetRecipeandMenu(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	id := c.Param("id")
	query := "SELECT recipe.id,recipe.description,menu.id,menu.name,menu.price FROM recipe INNER JOIN menu ON recipe.id = menu.idRecipeFK WHERE recipe.id = '" + id + "'"
	rows, err := db.Query(query)

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
		controller.SendRecipeSuccessResponse(c, responseRecipe)
	} else {
		responseRecipe.Message = "Recipe Get Failed"
		controller.SendRecipeErrorResponse(c, responseRecipe)
	}
}

//masih di usahakan :v
func RecordChangeMenu(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	var notchangeMenu modelMenu.Menu
	//If there is no changes
	foodId := c.Param("id")
	foodName := c.PostForm("name")
	foodPrice, _ := strconv.Atoi(c.PostForm("price"))

	if foodName == "" && foodPrice == foodPrice {
		foodName = notchangeMenu.Name
		foodPrice = int(notchangeMenu.Price)

		rows, err := db.Query("SELECT name,price FROM food")

		var notchangeMenus []modelMenu.Menu
		for rows.Next() {
			if err := rows.Scan(&notchangeMenu.Name, &notchangeMenu.Price); err != nil {
				log.Fatal(err)
			} else {
				notchangeMenus = append(notchangeMenus, notchangeMenu)
			}
		}

		var oldrepsonse modelMenu.MenuResponse
		if err == nil {
			oldrepsonse.Message = "Memunculkan Data yang tidak di update"
			oldrepsonse.Data = notchangeMenus
			controller.SendMenuSuccessResponse(c, oldrepsonse)
		} else {
			oldrepsonse.Message = "Data tidak ke load"
			controller.SendMenuErrorResponse(c, oldrepsonse)
		}
	} else {
		var changeMenus []modelMenu.Menu
		var changeMenu modelMenu.Menu
		_, query := db.Exec("UPDATE food SET name = ?,price = ? WHERE id = ?",
			foodName,
			foodPrice,
			foodId,
		)

		rows2, err2 := db.Query("SELECT name,price FROM food")
		for rows2.Next() {
			if err2 := rows2.Scan(&changeMenu.Name, &changeMenu.Price); err2 != nil {
				log.Fatal(err2)
			} else {
				changeMenus = append(changeMenus, changeMenu)
			}
		}

		var newresponse modelMenu.MenuResponse
		if query == nil && err2 == nil {
			newresponse.Message = "Update Food Success"
			newresponse.Data = changeMenus
			controller.SendMenuSuccessResponse(c, newresponse)
		} else {
			newresponse.Message = "Update Food Failed"
			controller.SendMenuErrorResponse(c, newresponse)
		}
	}
}
