package inventory

import (
	"fmt"
	"log"
	"net/http"
	model "github.com/AutoResto/material/entity"
	mtservice "github.com/AutoResto/material/service"
	"github.com/gin-gonic/gin"
)

//Insert Material
func AddNewMaterial(c *gin.Context) {
	c.Header("Content-Type", "Application/JSON")
	var response model.MaterialResponse
	errQuery := mtservice.NewMaterialRepository().InsertMaterialServiceDB(c)

	if errQuery == nil {
		response.Message = "Insert Material Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Insert Material Failed"
		c.JSON(http.StatusBadRequest, response)
	}

}
func ShowMaterialById(c *gin.Context){
	rows,err := mtservice.NewMaterialRepository().GetMaterialServiceById(c)

	if err != nil{
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material

	for rows.Next(){
		if err := rows.Scan(&material.Id,&material.Name,&material.Quantity,&material.Unit,&material.Inventory.Id);err != nil{
			log.Fatal(err.Error())
		}else{
			materials = append(materials, material)
		}
	}

	var response model.MaterialResponse

	if err == nil{
		response.Message = "Get Material Success"
		response.Data = materials
		c.JSON(http.StatusOK,response)
	}else{
		response.Message = "Get Material Failed"
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,response)
	}
}

//Get Material
func ShowMaterial(c *gin.Context) {
	rows, err := mtservice.NewMaterialRepository().GetMaterialServiceDB()
	if err != nil {
		log.Println(err)
	}

	var material model.Material
	var materials []model.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit,&material.Inventory.Id); err != nil {
			log.Fatal(err.Error())
		} else {
			materials = append(materials, material)
		}
	}

	var response model.MaterialResponse
	if err == nil {
		response.Message = "Get Material Success"
		response.Data = materials
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Get Material Failed"
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, response)
	}
}

//Update Material
func UpdateMaterial(c *gin.Context) {
	c.Header("Content-Type", "Application/JSON")
	errQuery := mtservice.NewMaterialRepository().UpdateMaterialServiceDB(c)
	var response model.MaterialResponse
	if errQuery == nil {
		response.Message = "Update Material Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Update Material Failed"
		c.JSON(http.StatusBadRequest, response)
	}
}

//Delete Material
func RemoveMaterial(c *gin.Context) {
	query := mtservice.NewMaterialRepository().DeleteMaterialServiceDB(c)
	var response model.MaterialResponse
	if query == nil {
		response.Message = "Delete Success"
		c.JSON(http.StatusOK, response)

	} else {
		response.Message = "Delete Failed"
		c.JSON(http.StatusBadRequest, response)
	}
}
