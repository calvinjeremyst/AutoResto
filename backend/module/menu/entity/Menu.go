package entity

type Menu struct {
	Id    int     `form:"id" json:"id"`
	Name  string  `form:"name" json:"name"`
	Price float64 `form:"price" json:"price"`
	
}

type MenuResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Menu `form:"data" json:"data"`
}

