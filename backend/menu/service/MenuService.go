package menu

import (
	"database/sql"
	"strconv"

	handler "github.com/AutoResto/handler"
	menu "github.com/AutoResto/menu/entity"
	mnrepo "github.com/AutoResto/menu/repository"
	"github.com/gin-gonic/gin"
)

type MenuRepo struct{
	mnrepo.MenuRepository
}

func NewMenuRepository() *MenuRepo{
	return &MenuRepo{}
}

//Select Menu Service
func (r *MenuRepo) SelectMenuServiceDB() (row *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()
	query := "SELECT id,name,price FROM menu"
	menu, err := db.Query(query)
	return menu, err
}

//Search Menu Service
func (r *MenuRepo) SearchMenuServiceDB(c *gin.Context) (row *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()

	menuName := c.Param("menu_name")

	query := "SELECT m.id, m.name, m.price FROM `menu` m WHERE m.name LIKE '" + menuName + "%'"

	menu, err := db.Query(query)
	return menu, err
}

//Get Availability Menu,utk cek apakah sudah ada / belum
func (r *MenuRepo) GetAvailabilityMenuServiceDB(c *gin.Context)(row *sql.Rows, err error){
	db := handler.Connect()
	defer db.Close()

	menuName := c.Param("name")
	query := "SELECT * FROM menu WHERE name = '"+menuName+"'"

	menu,err := db.Query(query)
	return menu,err
}

func (r *MenuRepo) UpdateMenuServiceDB(c *gin.Context) error {
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

func (r *MenuRepo) InsertMenuService(c *gin.Context) error {
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

func (r *MenuRepo) DeleteMenuServiceDB(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()
	menuId := c.Param("menu_id")

	_, errQuery := db.Exec("DELETE FROM menu WHERE id=?",
		menuId,
	)
	return errQuery
}
