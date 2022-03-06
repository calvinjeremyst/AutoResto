package entity

type Inventory struct {
	Id       int    `form:"id" json:"id"`
	Capacity int    `form:"capacity" json:"capacity"`
	Location string `form:"location" json:"location"`
}

type InventoryResponse struct {
	Message string      `form:"message" json:"message"`
	Data    []Inventory `form:"data" json:"data"`
}
