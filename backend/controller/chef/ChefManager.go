package chef

import (
	
	"log"
	"net/http"
	"strconv"

	controller "github.com/AutoResto/controller"
	modelMenu "github.com/AutoResto/module/menu/entity"
	modelRecipe "github.com/AutoResto/module/recipe/entity"
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

	var menu modelMenu.Menu
	var recipe modelRecipe.Recipe
	var foods []modelMenu.Menu
	var recipes []modelRecipe.Recipe

	for rows.Next(){
		if err := rows.Scan(&recipe.Id,&recipe.Description,&menu.Id,&menu.Name,&menu.Price);err != nil{
			log.Fatal(err)
		}else{
			recipes = append(recipes, recipe)
			foods = append(foods, menu)
			
		}
	}

	var responseFood modelMenu.MenuResponse
	var responseRecipe modelRecipe.RecipeResponse

	

	if err == nil{
		responseFood.Message = "Get Food Success"
		responseRecipe.Message = "Get Recipe Success"
		responseFood.Data = foods
		responseRecipe.Data = recipes
		sendFoodSuccessResponse(c,responseFood)
		sendRecipeSuccessResponse(c,responseRecipe)

	}else{
		responseFood.Message = "Food  Get Failed"
		responseRecipe.Message = "Recipe Get Failed"
		sendFoodErrorResponse(c,responseFood)
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


 