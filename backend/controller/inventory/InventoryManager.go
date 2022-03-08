package inventory

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AutoResto/controller"
	//controller "github.com/AutoResto/controller"
	model "github.com/AutoResto/module/material/entity"
	"github.com/gin-gonic/gin"
)


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
		sendMaterialSuccessResponse(c, response)
	} else {
		response.Message = "Insert Material Failed"
		sendMaterialErrorResponse(c, response)
	}

}

func SearchMaterial(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	namematerial := c.Param("name")

	query := "SELECT * FROM `material` WHERE name = '"+namematerial+"'"

	rows,err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material
	
	for rows.Next(){
		if err := rows.Scan(&material.Id,&material.Name,&material.Quantity,&material.Unit);err != nil{
			log.Fatal(err.Error())
		}else{
			materials = append(materials, material)
		}
	}
	
	var response model.MaterialResponse
	if err == nil {
		response.Message = "Search Material Success"
		response.Data = materials
		sendMaterialSuccessResponse(c,response)

	}else{
		response.Message = "Search material failed"
		fmt.Print(err)
		sendMaterialErrorResponse(c,response)
	}

}

func GetMaterial(c* gin.Context){
	db := controller.Connect()
	defer db.Close()

	query := "SELECT * FROM material"
	rows,err := db.Query(query)
	if err != nil{
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material

	for rows.Next(){
		if err := rows.Scan(&material.Id,&material.Name,&material.Quantity,&material.Unit);err != nil{
			log.Fatal(err.Error())
		}else{
			materials = append(materials, material)
		}
	}

	var response model.MaterialResponse
	if err == nil{
		response.Message = "Get Material Success"
		response.Data = materials
		sendMaterialSuccessResponse(c,response)
	}else{
		response.Message = "Get Material Failed"
		fmt.Println(err)
		sendMaterialErrorResponse(c,response)
	}
}

func UpdateMaterial(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	idMaterial := c.Param("id")
	materialName := c.PostForm("name")
	materialQuantity,_ := strconv.Atoi(c.PostForm("quantity"))
	materialUnit := c.PostForm("unit")

	rows, _ := db.Query("SELECT * fROM material WHERE id = '"+idMaterial+"'")

	var material model.Material
	for rows.Next(){
		if err := rows.Scan(&material.Id,&material.Name,&material.Quantity,&material.Unit);err!= nil{
			log.Fatal(err.Error())
		}
	}

	if materialName == ""{
		materialName = material.Name
	}

	if materialUnit == ""{
		materialUnit = material.Unit
	}

	_, errQuery := db.Exec("UPDATE material SET name = ?,quantity = ?,unit = ? WHERE id = ?",
				materialName,
				materialQuantity,
				materialUnit,
				idMaterial,
	)

	var response model.MaterialResponse
	if errQuery == nil{
		response.Message = "Update Material Success"
		sendMaterialSuccessResponse(c,response)
	}else{
		response.Message = "Update Material Failed"
		sendMaterialErrorResponse(c,response)
	}
}

func DeleteMaterial(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	idmaterial := c.Param("id")

	_,query := db.Exec("DELETE FROM material WHERE id = ?",idmaterial)

	var response model.MaterialResponse
	if query == nil{
		response.Message = "Delete Success"
		sendMaterialSuccessResponse(c,response)

	}else{
		response.Message = "Delete Failed"
		sendMaterialErrorResponse(c,response)
	}
}

func sendMaterialSuccessResponse(c *gin.Context, response model.MaterialResponse) {
	c.JSON(http.StatusOK, response)
}

func sendMaterialErrorResponse(c *gin.Context, response model.MaterialResponse) {
	c.JSON(http.StatusBadRequest, response)
}
