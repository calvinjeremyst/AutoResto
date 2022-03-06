package controller

import (
	"fmt"
	"log"

	entityMaterial "github.com/AutoResto/module/material/entity"
	"github.com/gin-gonic/gin"
)

// Get Material
func GetMaterial(c *gin.Context) {
	db := connect()
	defer db.Close()

	materialId := c.Param("material_id")

	query := "SELECT m.id, m.name, m.quantity, m.unit FROM material m WHERE m.id='" + materialId + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var material entityMaterial.Material
	var materials []entityMaterial.Material

	for rows.Next() {
		if err := rows.Scan(&material.Id, &material.Name, &material.Quantity, &material.Unit); err != nil {
			log.Fatal(err.Error())
			fmt.Println(material.Name)
		} else {
			materials = append(materials, material)
		}
	}

	var response entityMaterial.MaterialResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = materials
		sendMaterialSuccessresponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendMaterialErrorResponse(c, response)
	}
}
