package menu

import (
	"database/sql"
	"strconv"

	conn "github.com/AutoResto/dp/singleton"
	menu "github.com/AutoResto/menu/entity"
	mnrepo "github.com/AutoResto/menu/repository"
	recipe "github.com/AutoResto/recipe/entity"
	"github.com/gin-gonic/gin"
)

type MenuRepo struct {
	mnrepo.MenuRepository
}

func NewMenuRepository() *MenuRepo {
	return &MenuRepo{}
}

//Select Menu Service
func (r *MenuRepo) SelectMenuOwnerServiceDB() (row *sql.Rows, err error) {
	db := conn.Connect()
	defer db.Close()
	query := "SELECT menu.id,menu.name,menu.price,recipe.id,recipe.description FROM menu JOIN recipe ON recipe.id = menu.idRecipeFK"
	menu, err := db.Query(query)
	return menu, err
}

func (r *MenuRepo) SelectMenuChefServiceDB() (row *sql.Rows, err error) {

	db := conn.Connect()
	defer db.Close()

	query := "SELECT menu.id,menu.name,recipe.description FROM menu JOIN recipe ON recipe.id = menu.idRecipeFK"
	menu, err := db.Query(query)
	return menu, err
}

//Search Menu Service
func (r *MenuRepo) SearchMenuServiceDB(c *gin.Context) (row *sql.Rows, err error) {
	db := conn.Connect()
	defer db.Close()

	menuName := c.Param("menu_name")

	query := "SELECT m.id, m.name, m.price FROM `menu` m WHERE m.name LIKE '" + menuName + "%'"

	menu, err := db.Query(query)
	return menu, err
}

func (r *MenuRepo) GetMenuDetailByIdService(c *gin.Context) (row *sql.Rows, err error) {
	db := conn.Connect()
	defer db.Close()

	id := c.Param("id")
	query := "SELECT menu.name,menu.price,recipe.description FROM menu JOIN recipe ON recipe.id = menu.idRecipeFK WHERE recipe.id = '" + id + "'"
	menu, err := db.Query(query)
	return menu, err
}

//Get Availability Menu,utk cek apakah sudah ada / belum
func (r *MenuRepo) GetAvailabilityMenuServiceDB(c *gin.Context) (row *sql.Rows, err error) {
	db := conn.Connect()
	defer db.Close()

	menuName := c.Param("name")
	query := "SELECT * FROM menu WHERE name = '" + menuName + "'"

	menu, err := db.Query(query)
	return menu, err
}

func (r *MenuRepo) UpdateMenuServiceDB(c *gin.Context) error {
	db := conn.Connect()
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

func (r *MenuRepo) UpdateDetailMenuService(detailmenu recipe.Recipe, c *gin.Context) error {
	db := conn.Connect()
	defer db.Close()

	id := c.Param("id")

	_, query := db.Exec("UPDATE menu JOIN recipe ON recipe.id = menu.idRecipeFK SET menu.name = ?, menu.price = ?,recipe.description = ? WHERE recipe.id = ?",
		detailmenu.Menu.Name,
		detailmenu.Menu.Price,
		detailmenu.Description,
		id,
	)

	return query
}

func (r *MenuRepo) InsertMenuService(c *gin.Context) error {
	db := conn.Connect()
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
	db := conn.Connect()
	defer db.Close()

	menuId := c.Param("menu_id")

	_, errQuery := db.Exec("DELETE FROM recipe WHERE id=?",
		menuId,
	)

	return errQuery
}
