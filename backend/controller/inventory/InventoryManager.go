package inventory

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "github.com/AutoResto/module/material/entity"
	"github.com/gin-gonic/gin"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/autoresto?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db

}

func InsertMaterial(c *gin.Context) {

	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	unit := c.PostForm("unit")

	_, errQuery := db.Exec("INSERT INTO material(id,name,quantity,unit) VALUES(?,?,?,?)",
		id,
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
	db := connect()
	defer db.Close()

	idmaterial := c.Param("id")

	query := "SELECT * FROM `material` WHERE id = '"+idmaterial+"'"

	rows,err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material
	
	var id = material.GetId()
	var materialname = material.GetName()
	var materialquantity = material.GetQuantity()
	var materialunit = material.GetUnit()

	for rows.Next(){
		if err := rows.Scan(id,materialname,materialquantity,materialunit);err != nil{
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

	db := connect()
	defer db.Close()

	query := "SELECT * FROM material"
	rows,err := db.Query(query)
	if err != nil{
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material

	var idmaterial = material.GetId()
	var materialname = material.GetName()
	var materialquantity = material.GetQuantity()
	var materialunit = material.GetUnit()


	for rows.Next(){
		if err := rows.Scan(idmaterial,materialname,materialquantity,materialunit);err != nil{
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

	db := connect()
	defer db.Close()

	idMaterial := c.Param("id")
	materialName := c.PostForm("name")
	_,materialQuantity := strconv.Atoi(c.PostForm("quantity"))
	materialUnit := c.PostForm("unit")

	
	rows, _ := db.Query("SELECT * fROM material WHERE id = '"+idMaterial+"'")

	var material model.Material

	var id = material.GetId()
	var name = material.GetName()
	var qty = material.GetQuantity()
	var unt = material.GetUnit()


	for rows.Next(){
		if err := rows.Scan(id,name,qty,unt);err!= nil{
			log.Fatal(err.Error())
		}
	}

	if materialName == ""{
		materialName = name
	}

	if materialQuantity == materialQuantity{
		materialQuantity = materialQuantity 
	}

	if materialUnit == ""{
		materialUnit = unt
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

	db := connect()
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




