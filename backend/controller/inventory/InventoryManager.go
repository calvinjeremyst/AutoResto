package inventory

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AutoResto/controller"
	model "github.com/AutoResto/domain/material/entity"
	"github.com/gin-gonic/gin"
)

//Insert Material
func InsertMaterial(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	name := c.PostForm("name")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	unit := c.PostForm("unit")

	_, errQuery := db.Exec("INSERT INTO material(name,quantity,unit) VALUES(?,?,?)",

		name,
		quantity,
		unit)

	var response model.MaterialResponse

	if errQuery == nil {
		response.Message = "Insert Material Success"
		controller.SendMaterialSuccessResponse(c, response)
	} else {
		response.Message = "Insert Material Failed"
		controller.SendMaterialErrorResponse(c, response)
	}

}

//Get Material
func GetMaterial(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	query := "SELECT * FROM material"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}

	var response model.MaterialResponse
	if err == nil {
		response.Message = "Get Material Success"
		response.Data = materials
		controller.SendMaterialSuccessResponse(c, response)
	} else {
		response.Message = "Get Material Failed"
		fmt.Println(err)
		controller.SendMaterialErrorResponse(c, response)
	}
}

//Update Material
func UpdateMaterial(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	idMaterial := c.Param("id")
	materialName := c.PostForm("name")
	materialQuantity, _ := strconv.Atoi(c.PostForm("quantity"))
	materialUnit := c.PostForm("unit")

	rows, _ := db.Query("SELECT * fROM material WHERE id = '" + idMaterial + "'")

	var material model.Material
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

	var response model.MaterialResponse
	if errQuery == nil {
		response.Message = "Update Material Success"
		controller.SendMaterialSuccessResponse(c, response)
	} else {
		response.Message = "Update Material Failed"
		controller.SendMaterialErrorResponse(c, response)
	}
}

//Delete Material
func DeleteMaterial(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	idmaterial := c.Param("id")

	_, query := db.Exec("DELETE FROM material WHERE id = ?", idmaterial)

	var response model.MaterialResponse
	if query == nil {
		response.Message = "Delete Success"
		controller.SendMaterialSuccessResponse(c, response)

	} else {
		response.Message = "Delete Failed"
		controller.SendMaterialErrorResponse(c, response)
	}
}
