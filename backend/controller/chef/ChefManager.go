package chef

import (
	
	"log"
	"net/http"
	"strconv"

	controller "github.com/AutoResto/controller"
	modelRecipe "github.com/AutoResto/module/recipe/entity"
	modelMenu "github.com/AutoResto/module/menu/entity"
	"github.com/gin-gonic/gin"
)

func GetRecipeandMenu(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

  id := c.Param("id")
	query := "SELECT recipe.id,recipe.description,menu.id,menu.name,menu.price FROM recipe INNER JOIN menu ON recipe.id = menu.idRecipeFK WHERE recipe.id = '"+id+"'"
	rows,err := db.Query(query)
	
	if err != nil{
		log.Println(err)
	}

	var recipe modelRecipe.Recipe
	var recipes []modelRecipe.Recipe
	for rows.Next(){
		if err := rows.Scan(&recipe.Id,&recipe.Description,&recipe.Menu.Id,&recipe.Menu.Name,&recipe.Menu.Price);err != nil{
			log.Fatal(err)
		}else{
			recipes = append(recipes, recipe)
		}
	}
	
	var responseRecipe modelRecipe.RecipeResponse
	if err == nil{
		responseRecipe.Message = "Get Recipe Success"
		responseRecipe.Data = recipes
		sendRecipeSuccessResponse(c,responseRecipe)
	}else{
		responseRecipe.Message = "Recipe Get Failed"
		sendRecipeErrorResponse(c,responseRecipe)
	}
 }

 //masih di usahakan :v
 func RecordChangeMenu(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	var notchangeMenu modelMenu.Menu
	//If there is no changes
	foodId := c.Param("id")
	foodName := c.PostForm("name")
	foodPrice,_ := strconv.Atoi(c.PostForm("price"))

	if foodName == "" && foodPrice == foodPrice{
		foodName = notchangeMenu.Name
		foodPrice = int(notchangeMenu.Price)

		rows,err := db.Query("SELECT name,price FROM food")
		
    var notchangeMenus 	  []modelMenu.Menu
		for rows.Next(){
			if err := rows.Scan(&notchangeMenu.Name,&notchangeMenu.Price);err != nil{
				log.Fatal(err)
			}else{
				notchangeMenus = append(notchangeMenus,notchangeMenu)
			}		
		}
		
    var oldrepsonse modelMenu.MenuResponse
		if err == nil{
			oldrepsonse.Message = "Memunculkan Data yang tidak di update"
			oldrepsonse.Data = notchangeMenus
			sendFoodSuccessResponse(c,oldrepsonse)
		}else{
			oldrepsonse.Message = "Data tidak ke load"
			sendFoodSuccessResponse(c,oldrepsonse)
		}
	}else{
		var changeMenus []modelMenu.Menu
		var changeMenu modelMenu.Menu
		_, query := db.Exec("UPDATE food SET name = ?,price = ? WHERE id = ?",
			foodName,
			foodPrice,
			foodId,
		)
    
		rows2,err2 := db.Query("SELECT name,price FROM food")
		for rows2.Next(){
			if err2 := rows2.Scan(&changeMenu.Name,&changeMenu.Price);err2 != nil{
				log.Fatal(err2)
			}else{
				changeMenus = append(changeMenus, changeMenu)
			}
		}

		var newresponse modelMenu.MenuResponse
		if query == nil && err2 == nil {
			newresponse.Message = "Update Food Success"
			newresponse.Data = changeMenus
			sendFoodSuccessResponse(c,newresponse)
		}else{
			newresponse.Message = "Update Food Failed"
			sendFoodErrorResponse(c,newresponse)
		}
	}	
 }

 func sendFoodSuccessResponse(c *gin.Context, response modelMenu.MenuResponse){
	c.JSON(http.StatusOK, response)
 }

 func sendRecipeSuccessResponse(c *gin.Context, response modelRecipe.RecipeResponse){
	c.JSON(http.StatusOK, response)
 }

 func sendFoodErrorResponse(c *gin.Context,response modelMenu.MenuResponse){
	 c.JSON(http.StatusBadRequest,response)
 }

 func sendRecipeErrorResponse(c *gin.Context,response modelRecipe.RecipeResponse){
	c.JSON(http.StatusBadRequest,response)
 }
