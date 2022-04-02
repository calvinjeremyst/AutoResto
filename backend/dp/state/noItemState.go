package state

import (
	iv "github.com/AutoResto/inventory/controller/inventory"
	modelMaterial "github.com/AutoResto/material/entity"
	"github.com/gin-gonic/gin"
)

type noItemState struct{
	checker *checker
}

func (n *noItemState) requestItem(){
	
	var response modelMaterial.MaterialResponse
	response.Message = "Material Not Available"
}

func (n *noItemState) addItem(c *gin.Context){
	iv.AddNewMaterial(c)
}

func (n *noItemState) dispenseItem(){
	var response modelMaterial.MaterialResponse
	response.Message = "Material Not Available"
}