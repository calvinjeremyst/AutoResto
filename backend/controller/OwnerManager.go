package controller

import (
	"log"
	"net/http"
	"strconv"

	model "github.com/AutoResto/module/menu/entity"

	"github.com/gin-gonic/gin"
)

// Get All Menus
func GetAllMenus(c *gin.Context) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM menu"

	nama := c.Query("name")
	if nama != "" {
		query += " WHERE name='" + nama + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var menu model.Menu
	var menus []model.Menu
	for rows.Next() {
		if err := rows.Scan(&menu.Id, &menu.Name, &menu.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			menus = append(menus, menu)
		}
	}

	var response model.MenuResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = menus
		sendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendMenuErrorResponse(c, response)
	}
}

// Insert Menu
func InsertMenu(c *gin.Context) {
	db := Connect()
	defer db.Close()

	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))

	_, errQuery := db.Exec("INSERT INTO menu(name, price) values (?,?)",
		name,
		price,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Insert Menu Success"
		sendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Insert Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}

// Delete Menu
func DeleteMenu(c *gin.Context) {
	db := Connect()
	defer db.Close()

	menuId := c.Param("menu_id")

	_, errQuery := db.Exec("DELETE FROM menu WHERE id=?",
		menuId,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Delete Menu Success"
		sendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Delete Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}

func sendMenuSuccessResponse(c *gin.Context, response model.MenuResponse) {
	c.JSON(http.StatusOK, response)
}

func sendMenuErrorResponse(c *gin.Context, response model.MenuResponse) {
	c.JSON(http.StatusBadRequest, response)
}
