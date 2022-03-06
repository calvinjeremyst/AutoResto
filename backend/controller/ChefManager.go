package controller

import (
	"log"
	"strconv"

	model "github.com/AutoResto/module/menu/entity"

	"github.com/gin-gonic/gin"
)

// Update Menu
func UpdateMenu(c *gin.Context) {
	db := Connect()
	defer db.Close()

	menuId := c.Param("id")
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))

	rows, _ := db.Query("SELECT * FROM menu WHERE id='" + menuId + "'")
	var menu model.Menu
	for rows.Next() {
		if err := rows.Scan(&menu.Id, &menu.Name, &menu.Price); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if name == "" {
		name = menu.Name
	}

	if price == 0 {
		price = int(menu.Price)
	}

	_, errQuery := db.Exec("UPDATE menu SET name = ?, price = ?, WHERE id=?",
		name,
		price,
		menuId,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Update Menu Success"
		sendMenuSuccessResponse(c, response)
	} else {
		response.Message = "Update Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}
