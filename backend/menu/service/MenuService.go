package menu

import (
	"database/sql"
	"strconv"

	handler "github.com/AutoResto/handler"
	menu "github.com/AutoResto/menu/entity"
	"github.com/gin-gonic/gin"
)

//Select Menu Service
func SelectMenuServiceDB()(row *sql.Rows,err error){
	db := handler.Connect()
	defer db.Close()
	query := "SELECT id,name,price FROM menu"
	menu, err := db.Query(query)
	return menu,err
}

//Search Menu Service
func SearchMenuServiceDB(c *gin.Context)(row *sql.Rows,err error){
	db := handler.Connect()
	defer db.Close()

	menuName := c.Param("menu_name")

	query := "SELECT m.id, m.name, m.price FROM `menu` m WHERE m.name LIKE '"+menuName+"%'"

	menu,err := db.Query(query)
	return menu,err
}

func UpdateMenuServiceDB(c* gin.Context) error{
	db := handler.Connect()
	defer db.Close()

	menuId := c.Param("menu_id")
	name := c.PostForm("name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 32)
	var menu menu.Menu
		// Jika kosong dimasukkan nilai lama
		if name == "" {
			name = menu.Name
		}
	
		if price == 0 {
			price = menu.Price
		}
		_, errQuery := db.Exec("UPDATE menu SET name = ?, price = ? WHERE id=?",
			name,
			price,
			menuId,
		)
		
		return errQuery
}

func InsertMenuService(c* gin.Context) error{
	db := handler.Connect()
	defer db.Close()
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	description := c.PostForm("description")

	_, errQuery := db.Exec("INSERT INTO menu(name,price,idRecipeFK) VALUES (?,?,(SELECT id from recipe where recipe.description='"+description+"'));",
		name,
		price,
	)

	return errQuery
}

func DeleteMenuServiceDB(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()
	menuId := c.Param("menu_id")

	_, errQuery := db.Exec("DELETE FROM menu WHERE id=?",
		menuId,
	)
	return errQuery
}