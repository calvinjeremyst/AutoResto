package material

import (
	"database/sql"
	"log"
	"strconv"
	//"strings"

	handler "github.com/AutoResto/handler"
	material "github.com/AutoResto/material/entity"
	"github.com/gin-gonic/gin"
)

func GetMaterialServiceDB() (rows *sql.Rows, err error) {
	db := handler.Connect()
	defer db.Close()

	query := "SELECT * FROM material"
	material, err := db.Query(query)

	return material, err

}

func GetMaterialServiceById(c *gin.Context)(rows *sql.Rows,err error){

	db := handler.Connect()
	defer db.Close()
	id := c.Param("id")
	query := "SELECT * FROM material WHERE id = '"+id+"'"

	material,err := db.Query(query)
	return material,err

}

func SearchMaterialServiceDB(c *gin.Context) (rows *sql.Rows, err error) {
	c.Header("Content-Type", "Application/JSON")
	db := handler.Connect()
	defer db.Close()

	materialName := c.Param("material_name")

	query := "SELECT m.id, m.name, m.quantity, m.unit FROM material m WHERE m.name LIKE '" + materialName + "%'"
	material, err := db.Query(query)
	return material, err

}

func InsertMaterialHelperServiceDB(c *gin.Context) error {

	db := handler.Connect()
	defer db.Close()

	materialName := c.PostForm("material")
	//materialArr := strings.Split(materialName, ",")
	quantityMaterial,_ := strconv.Atoi(c.PostForm("quantity_material"))
	unitMaterial := c.PostForm("unit_material")

	_, errQuery := db.Exec("INSERT INTO material(name,quantity,unit) VALUES(?,?,?)",
		materialName,
		quantityMaterial,
		unitMaterial,
	)
	return errQuery
}

func InsertMaterialServiceDB(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()

	materialname := c.PostForm("name")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	unit := c.PostForm("unit")

	_, errQuery := db.Exec("INSERT INTO material(name,quantity,unit) VALUES(?,?,?)", materialname, quantity, unit)
	return errQuery

}

func UpdateMaterialServiceDB(c *gin.Context) error {
	db := handler.Connect()
	defer db.Close()

	idMaterial := c.Param("id")
	materialName := c.PostForm("name")
	materialQuantity, _ := strconv.Atoi(c.PostForm("quantity"))
	materialUnit := c.PostForm("unit")

	rows, _ := db.Query("SELECT * fROM material WHERE id = '" + idMaterial + "'")

	var material material.Material
	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit); err != nil {
			log.Fatal(err.Error())
		}
	}
	_, errQuery := db.Exec("UPDATE material SET name = ?,quantity = ?,unit = ? WHERE id = ?",
		materialName,
		materialQuantity,
		materialUnit,
		idMaterial,
	)
	return errQuery
}

func DeleteMaterialServiceDB(c *gin.Context) error {

	db := handler.Connect()
	defer db.Close()
	idmaterial := c.Param("id")
	_, query := db.Exec("DELETE FROM material WHERE id = ?", idmaterial)
	return query

}
