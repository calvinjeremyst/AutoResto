package entity

type Food struct {
	Id    int     `form:"id" json:"id"`
	Name  string  `form:"name" json:"name"`
	Price float64 `form:"price" json:"price"`
	
}

type FoodResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Food `form:"data" json:"data"`
}

