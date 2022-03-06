package entity

type Material struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Quantity int    `form:"quantity" json:"quantity"`
	Unit 	 string `form : "unit" json : "unit"`
}

type MaterialResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Material `form:"data" json:"data"`
}



