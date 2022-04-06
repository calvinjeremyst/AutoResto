package entity
import (

	modelInv "github.com/AutoResto/inventory/entity"
)

type Material struct {
	Id       	  int    `form : "id" json:"id"`
	Name     	 string `form : "name" json:"name"`
	Quantity 	 int    `form : "quantity" json:"quantity"`
	Unit     	 string `form : "unit" json:"unit"`
	Inventory modelInv.Inventory `form : "id_inventory" json : "id_inventory"`
}
