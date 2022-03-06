package chef

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/AutoResto/controller"
	modelFood "github.com/AutoResto/module/food/entity"
	modelRecipe "github.com/AutoResto/module/recipe/entity"
	"github.com/gin-gonic/gin"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/autoresto?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db

}


 func GetRecipeandMenu(c *gin.Context){

	db := connect()
	defer db.Close()

	query := "SELECT food.name AS foodname recipe.description AS recipe FROM food JOIN recipe ON food.id = recipe.idFood"
	rows,err := db.Query(query)
	
	if err != nil{
		log.Println(err)
	}

	var food modelFood.Food
	var recipe modelRecipe.Recipe
	var foods []modelFood.Food
	var recipes []modelRecipe.Recipe

	for rows.Next(){
		if err := rows.Scan(food.Name,recipe.Description);err != nil{
			log.Fatal(err)
		}else{
			foods = append(foods, food)
			recipes = append(recipes, recipe)
		}
	}

	var responseFood modelFood.FoodResponse
	var responseRecipe modelRecipe.RecipeResponse

	

	if err == nil{
		responseFood.Message = "Get Food Success"
		responseRecipe.Message = "Get Recipe Success"
		responseFood.Data = foods
		responseRecipe.Data = recipes
		sendFoodSuccessResponse(c,responseFood)
		sendRecipeSuccessResponse(c,responseRecipe)

	}else{
		responseFood.Message = "Food  Get Success"
		responseRecipe.Message = "Recipe Get Success"
		sendFoodErrorResponse(c,responseFood)
		sendRecipeErrorResponse(c,responseRecipe)
	}

 }


 func RecordChangeMenu(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	var notchangeMenu modelFood.Food
	//If there is no changes
	foodId := c.Param("id")
	foodName := c.PostForm("name")
	foodPrice,_ := strconv.Atoi(c.PostForm("price"))
	

	if foodName == "" && foodPrice == foodPrice{
		foodName = notchangeMenu.Name
		foodPrice = int(notchangeMenu.Price)

		rows,err := db.Query("SELECT name,price FROM food")
		var notchangeMenus 	  []modelFood.Food
		

		for rows.Next(){
			if err := rows.Scan(&notchangeMenu.Name,&notchangeMenu.Price);err != nil{
				log.Fatal(err)
			}else{
				notchangeMenus = append(notchangeMenus,notchangeMenu)
			}	
			
		}
		var oldrepsonse modelFood.FoodResponse
		if err == nil{
			oldrepsonse.Message = "Memunculkan Data yang tidak di update"
			oldrepsonse.Data = notchangeMenus
			sendFoodSuccessResponse(c,oldrepsonse)
		}else{
			oldrepsonse.Message = "Data tidak ke load"
			sendFoodSuccessResponse(c,oldrepsonse)
		}
	}else{
		var changeMenus 	  []modelFood.Food
		var changeMenu modelFood.Food
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



		var newresponse modelFood.FoodResponse
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

 func sendFoodSuccessResponse(c *gin.Context, response modelFood.FoodResponse){
	c.JSON(http.StatusOK, response)
 }

 func sendRecipeSuccessResponse(c *gin.Context, response modelRecipe.RecipeResponse){
	c.JSON(http.StatusOK, response)
 }

 func sendFoodErrorResponse(c *gin.Context,response modelFood.FoodResponse){
	 c.JSON(http.StatusBadRequest,response)

 }

 func sendRecipeErrorResponse(c *gin.Context,response modelRecipe.RecipeResponse){
	c.JSON(http.StatusBadRequest,response)

}


 